package pg

import (
	"context"
	"fmt"
	"net/http"
	"sort"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/splathon/serror"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
	"golang.org/x/sync/errgroup"
)

func (h *Handler) GetResult(ctx context.Context, params result.GetResultParams) (*models.Results, error) {
	var eg errgroup.Group

	var (
		qualifiers  []*Qualifier
		tournaments []*Tournament
		qmatches    []*Match // Matches in Qualifiers.
		tmatches    []*Match // Matches in Tournament.
		teams       []*Team
		rooms       []*Room
	)

	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("round asc").Find(&qualifiers).Error
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("round asc").Find(&tournaments).Error
	})

	eg.Go(func() error {
		qids := h.db.Table("qualifiers").Select("id").Where("event_id = ?", eventID).QueryExpr()
		query := h.db.Where("qualifier_id in (?)", qids)
		if params.TeamID != nil {
			teamID := *params.TeamID
			query = h.db.Where("qualifier_id in (?) AND (team_id = ? OR opponent_id = ?)", qids, teamID, teamID)
		}
		return query.Find(&qmatches).Error
	})

	eg.Go(func() error {
		tids := h.db.Table("tournaments").Select("id").Where("event_id = ?", eventID).QueryExpr()
		query := h.db.Where("tournament_id in (?)", tids)
		if params.TeamID != nil {
			teamID := *params.TeamID
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

	if err := checkTeamFound(params, teams); err != nil {
		return nil, err
	}

	return buildResult(qualifiers, tournaments, qmatches, tmatches, teams, rooms), nil
}

func checkTeamFound(params result.GetResultParams, teams []*Team) error {
	if params.TeamID == nil {
		return nil
	}
	for _, t := range teams {
		if t.Id == *params.TeamID {
			return nil
		}
	}
	return &serror.Error{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("team_id=%d not found", *params.TeamID),
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
