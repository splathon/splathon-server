package pg

import (
	"context"
	"fmt"
	"sort"

	"github.com/go-openapi/swag"
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
		if err := h.db.Where("event_id = ?", eventID).Order("round asc").Find(&qualifiers).Error; err != nil {
			return err
		}
		qids := make([]int64, 0, len(qualifiers))
		for _, q := range qualifiers {
			qids = append(qids, q.Id)
		}
		return h.db.Where("qualifier_id in (?)", qids).Find(&matches).Error
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

	return buildResult(qualifiers, matches, teams, rooms), nil
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
			Name:  swag.String(fmt.Sprintf("予選第%dラウンド", q.Round)),
			Round: q.Round,
			Rooms: make([]*models.Room, 0, len(rooms)),
		}

		for _, r := range rooms {
			room := &models.Room{
				ID:      int32(r.Id),
				Name:    swag.String(r.Name),
				Matches: make([]*models.Match, 0, len(ms[q.Id][r.Id])),
			}

			for _, m := range ms[q.Id][r.Id] {
				room.Matches = append(room.Matches, convertMatch(m, teamMap))
				sort.Slice(room.Matches, func(i, j int) bool {
					return room.Matches[i].Order < room.Matches[j].Order
				})
			}

			round.Rooms = append(round.Rooms, room)
		}

		result.Qualifiers = append(result.Qualifiers, round)
	}
	return result
}

func convertMatch(m *Match, teamMap map[int64]*Team) *models.Match {
	match := &models.Match{
		ID:    swag.Int32(int32(m.Id)),
		Order: int32(m.Order),
	}

	if t, ok := teamMap[m.TeamId]; ok {
		match.TeamAlpha = convertTeam(t)
	}
	if t, ok := teamMap[m.OpponentId]; ok {
		match.TeamBravo = convertTeam(t)
	}

	if m.TeamPoints > 0 && m.OpponentPoints > 0 && m.TeamPoints == m.OpponentPoints {
		match.Winner = models.MatchWinnerDraw
	} else if m.TeamPoints > m.OpponentPoints {
		match.Winner = models.MatchWinnerAlpha
	} else if m.TeamPoints < m.OpponentPoints {
		match.Winner = models.MatchWinnerBravo
	}
	return match
}

func convertTeam(t *Team) *models.Team {
	return &models.Team{
		ID:          swag.Int32(int32(t.Id)),
		CompanyName: swag.String(t.CompanyName),
		Name:        swag.String(t.Name),
	}
}
