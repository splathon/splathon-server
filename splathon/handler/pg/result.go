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

	var qualifiers []*Qualifier
	var matches []*Match
	var teams []*Team
	var rooms []*Room

	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("round asc").Find(&qualifiers).Error
	})

	eg.Go(func() error {
		qids := h.db.Table("qualifiers").Select("id").Where("event_id = ?", eventID).QueryExpr()
		query := h.db.Where("qualifier_id in (?)", qids)
		if params.TeamID != nil {
			teamID := *params.TeamID
			query = h.db.Where("qualifier_id in (?) AND (team_id = ? OR opponent_id = ?)", qids, teamID, teamID)
		}
		return query.Find(&matches).Error
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

	return buildResult(qualifiers, matches, teams, rooms), nil
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

func buildResult(qualifiers []*Qualifier, matches []*Match, teams []*Team, rooms []*Room) *models.Results {
	// qualifier_id => room_id => Match
	ms := make(map[int64]map[int64][]*Match)
	for _, m := range matches {
		if _, ok := ms[m.QualifierId]; !ok {
			ms[m.QualifierId] = make(map[int64][]*Match)
		}
		if _, ok := ms[m.QualifierId][m.RoomId]; !ok {
			ms[m.QualifierId][m.RoomId] = make([]*Match, 0)
		}
		ms[m.QualifierId][m.RoomId] = append(ms[m.QualifierId][m.RoomId], m)
	}

	// Past splathon may not have room.
	if len(rooms) == 0 {
		rooms = append(rooms, &Room{Name: "Unknown"})
	}

	// team_id => Team
	teamMap := make(map[int64]*Team)
	for _, t := range teams {
		teamMap[t.Id] = t
	}

	result := &models.Results{
		Qualifiers: make([]*models.Round, 0, len(qualifiers)),
	}
	for _, q := range qualifiers {
		round := &models.Round{
			Name:  swag.String(qualifierRoundName(int(q.Round))),
			Round: q.Round,
			Rooms: make([]*models.Room, 0, len(rooms)),
		}

		for _, r := range rooms {
			addRoom := false
			room := &models.Room{
				ID:      int32(r.Id),
				Name:    swag.String(r.Name),
				Matches: make([]*models.Match, 0, len(ms[q.Id][r.Id])),
			}

			for _, m := range ms[q.Id][r.Id] {
				addRoom = true
				room.Matches = append(room.Matches, convertMatch(m, teamMap))
				sort.Slice(room.Matches, func(i, j int) bool {
					return room.Matches[i].Order < room.Matches[j].Order
				})
			}

			if addRoom {
				round.Rooms = append(round.Rooms, room)
			}
		}

		result.Qualifiers = append(result.Qualifiers, round)
	}
	return result
}
