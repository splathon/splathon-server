package pg

import (
	"context"
	"sort"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"golang.org/x/sync/errgroup"
)

const qualifierMaxBattleNum = 2

func (h *Handler) GetMatch(ctx context.Context, params match.GetMatchParams) (*models.Match, error) {
	var eg errgroup.Group

	var (
		match   Match
		teams   []*Team
		battles []*Battle
	)

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

	seenBattleOrders := make(map[int]bool)
	for _, b := range battles {
		m.Battles = append(m.Battles, convertBattle(b, m))
		seenBattleOrders[int(b.Order)] = true
	}

	// Fill in not-finished battles.
	// TODO(haya14busa): register and get theses magic numbers from database.
	maxBattleNum := qualifierMaxBattleNum
	if match.QualifierId == 0 {
		maxBattleNum = 3
	}
	for order := 1; order <= maxBattleNum; order++ {
		if seenBattleOrders[order] {
			continue
		}
		m.Battles = append(m.Battles, &models.Battle{Order: swag.Int32(int32(order))})
	}
	sort.Slice(m.Battles, func(i, j int) bool {
		return *m.Battles[i].Order < *m.Battles[j].Order
	})
	return m, nil
}
