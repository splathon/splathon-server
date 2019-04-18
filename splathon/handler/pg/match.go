package pg

import (
	"context"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"golang.org/x/sync/errgroup"
)

func (h *Handler) GetMatch(ctx context.Context, params match.GetMatchParams) (*models.Match, error) {
	var eg errgroup.Group

	var (
		eventID int64
		match   Match
		teams   []*Team
		battles []*Battle
	)

	eg.Go(func() error {
		var err error
		eventID, err = h.queryInternalEventID(params.EventID)
		return err
	})

	eg.Go(func() error {
		if err := h.db.Where("id = ?", params.MatchID).Find(&match).Error; err != nil {
			return err
		}
		return h.db.Where("id = ? OR id = ?", match.TeamId, match.OpponentId).Find(&teams).Error
	})

	eg.Go(func() error {
		return h.db.Where("match_id = ?", params.MatchID).Order(`"order" asc`).Find(&battles).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	// team_id => Team
	teamMap := make(map[int64]*Team)
	for _, t := range teams {
		teamMap[t.Id] = t
	}
	m := convertMatch(&match, teamMap)
	for _, b := range battles {
		m.Battles = append(m.Battles, convertBattle(b, m))
	}
	return m, nil
}
