package pg

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/splathon/serror"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
	"golang.org/x/sync/errgroup"
)

func (h *Handler) GetResult(ctx context.Context, params result.GetResultParams) (*models.Results, error) {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}
	isAdmin := params.XSPLATHONAPITOKEN != nil && h.checkAdminAuth(*params.XSPLATHONAPITOKEN) == nil
	var teamID int64
	if params.TeamID != nil {
		teamID = *params.TeamID
	}

	// Do not use cache for requests from admin because it needs latest results
	// regardless of release flag.
	shouldUseCache := !isAdmin && teamID == 0
	if shouldUseCache {
		if ok, result := h.maybeGetResultFromCache(eventID); ok {
			return result, nil
		}
	}

	key := fmt.Sprintf("GetResult,event_id=%v,admin=%v,team_id=%v", eventID, isAdmin, teamID)
	v, err, shared := h.sfgroup.Do(key, func() (interface{}, error) {
		return h.getResultInternal(ctx, eventID, isAdmin, teamID)
	})
	if err != nil {
		return nil, err
	}
	if shared {
		log.Println("[INFO] get result from shared data")
	}
	result := v.(*models.Results)

	// Cache result.
	if shouldUseCache {
		h.resultCacheMu.Lock()
		defer h.resultCacheMu.Unlock()
		h.resultCache[eventID] = &resultCache{
			result:    result,
			timestamp: time.Now(),
		}
	}
	return result, nil
}

func (h *Handler) maybeGetResultFromCache(eventID int64) (bool, *models.Results) {
	h.resultCacheMu.Lock()
	defer h.resultCacheMu.Unlock()
	if cache, ok := h.resultCache[eventID]; ok && time.Now().Sub(cache.timestamp) < 1*time.Minute {
		log.Println("[INFO] get result from cache")
		return true, cache.result
	}
	return false, nil
}

func (h *Handler) getResultInternal(ctx context.Context, eventID int64,
	ignoreReleaseFlag bool, teamID int64) (*models.Results, error) {
	var eg errgroup.Group

	var (
		qualifiers  []*Qualifier
		tournaments []*Tournament
		qmatches    []*Match // Matches in Qualifiers.
		tmatches    []*Match // Matches in Tournament.
		teams       []*Team
		rooms       []*Room
	)

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("round asc").Find(&qualifiers).Error
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("round asc").Find(&tournaments).Error
	})

	eg.Go(func() error {
		qids := h.db.Table("qualifiers").Select("id").Where("event_id = ?", eventID).QueryExpr()
		query := h.db.Where("qualifier_id in (?)", qids)
		if teamID != 0 {
			query = h.db.Where("qualifier_id in (?) AND (team_id = ? OR opponent_id = ?)", qids, teamID, teamID)
		}
		return query.Find(&qmatches).Error
	})

	eg.Go(func() error {
		tids := h.db.Table("tournaments").Select("id").Where("event_id = ?", eventID).QueryExpr()
		query := h.db.Where("tournament_id in (?)", tids)
		if teamID != 0 {
			query = h.db.Where("tournament_id in (?) AND (team_id = ? OR opponent_id = ?)", tids, teamID, teamID)
		}
		return query.Find(&tmatches).Error
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Find(&teams).Error
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("id asc").Find(&rooms).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	if err := checkTeamFound(teamID, teams); err != nil {
		return nil, err
	}

	releasedRound, err := GetQualifierRelease(ctx, eventID)
	if err != nil {
		return nil, err
	}
	var qs []*Qualifier
	for _, q := range qualifiers {
		if ignoreReleaseFlag || releasedRound == -1 || q.Round <= releasedRound {
			qs = append(qs, q)
		}
	}
	return buildResult(qs, tournaments, qmatches, tmatches, teams, rooms), nil
}

func checkTeamFound(teamID int64, teams []*Team) error {
	if teamID == 0 {
		return nil
	}
	for _, t := range teams {
		if t.Id == teamID {
			return nil
		}
	}
	return &serror.Error{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("team_id=%d not found", teamID),
	}
}

// qualifier/tournament id => room_id => Match
type matchMap map[int64]map[int64][]*Match

func buildMatchMap(matches []*Match, idFunc func(m *Match) int64) matchMap {
	ms := make(map[int64]map[int64][]*Match)
	for _, m := range matches {
		roundID := idFunc(m)
		if _, ok := ms[roundID]; !ok {
			ms[roundID] = make(map[int64][]*Match)
		}
		if _, ok := ms[roundID][m.RoomId]; !ok {
			ms[roundID][m.RoomId] = make([]*Match, 0)
		}
		ms[roundID][m.RoomId] = append(ms[roundID][m.RoomId], m)
	}
	return ms
}

func buildResult(qualifiers []*Qualifier, tournaments []*Tournament, qmatches []*Match, tmatches []*Match, teams []*Team, rooms []*Room) *models.Results {
	// qualifier_id => room_id => Match
	qms := buildMatchMap(qmatches, func(m *Match) int64 { return m.QualifierId.Int64 })
	// tournament_id => room_id => Match
	tms := buildMatchMap(tmatches, func(m *Match) int64 { return m.TournamentId.Int64 })

	// Past splathon may not have room.
	if len(rooms) == 0 {
		rooms = append(rooms, &Room{Name: "Unknown"})
	}

	result := &models.Results{
		Qualifiers: make([]*models.Round, 0, len(qualifiers)),
	}
	for _, q := range qualifiers {
		if r, ok := buildRound(q, rooms, teams, qms); ok {
			result.Qualifiers = append(result.Qualifiers, r)
		}
	}
	ts := make([]*models.Round, 0, len(tournaments))
	for _, t := range tournaments {
		if r, ok := buildRound(t, rooms, teams, tms); ok {
			ts = append(ts, r)
		}
	}
	if len(ts) > 0 {
		result.Tournament = ts
	}
	return result
}

func buildRound(in Round, rooms []*Room, teams []*Team, mmap matchMap) (*models.Round, bool) {
	ok := false
	round := &models.Round{
		Name:  swag.String(in.GetName()),
		Round: in.GetRoundNumber(),
		Rooms: make([]*models.Room, 0, len(rooms)),
	}

	// team_id => Team
	teamMap := make(map[int64]*Team)
	for _, t := range teams {
		teamMap[t.Id] = t
	}

	for _, r := range rooms {
		addRoom := false
		room := &models.Room{
			ID:      int32(r.Id),
			Name:    swag.String(r.Name),
			Matches: make([]*models.Match, 0, len(mmap[in.GetID()][r.Id])),
		}

		for _, m := range mmap[in.GetID()][r.Id] {
			addRoom = true
			ok = true
			room.Matches = append(room.Matches, convertMatch(m, teamMap))
			sort.Slice(room.Matches, func(i, j int) bool {
				return room.Matches[i].Order < room.Matches[j].Order
			})
		}

		if addRoom {
			round.Rooms = append(round.Rooms, room)
		}
	}
	return round, ok
}
