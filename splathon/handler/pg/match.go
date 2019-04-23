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

func getMaxBattleNum(m Match) int {
	// TODO(haya14busa): register and get theses magic numbers from database.
	n := qualifierMaxBattleNum
	if !m.QualifierId.Valid {
		n = 3
	}
	return n
}

func (h *Handler) GetMatch(ctx context.Context, params match.GetMatchParams) (*models.Match, error) {
	var eg errgroup.Group

	var (
		match        Match
		teams        []*Team
		participants []*Participant
		battles      []*Battle
		roundName    string
	)

	eg.Go(func() error {
		if err := h.db.Where("id = ?", params.MatchID).Find(&match).Error; err != nil {
			return err
		}

		// Fetch team.
		eg.Go(func() error {
			return h.db.Where("id = ? OR id = ?", match.TeamId, match.OpponentId).Find(&teams).Error
		})

		// Fetch participants.
		eg.Go(func() error {
			return h.db.Where("team_id = ? OR team_id = ?", match.TeamId, match.OpponentId).Order("id asc").Find(&participants).Error
		})

		// Fetch round name.
		eg.Go(func() error {
			if match.QualifierId.Valid {
				// var round int
				var q Qualifier
				if err := h.db.Select("round").Where("id = ?", match.QualifierId).Find(&q).Error; err != nil {
					return err
				}
				roundName = q.GetName()
			} else if match.TournamentId.Valid {
				var t Tournament
				if err := h.db.Select("name").Where("id = ?", match.TournamentId).Find(&t).Error; err != nil {
					return err
				}
				roundName = t.GetName()
			}
			return nil
		})

		return nil
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

	// Fill in members.
	m.TeamAlpha.Members = make([]*models.Member, 0)
	m.TeamBravo.Members = make([]*models.Member, 0)
	for _, p := range participants {
		var t *models.Team
		switch p.TeamId.Int64 {
		case match.TeamId:
			t = m.TeamAlpha
		case match.OpponentId:
			t = m.TeamBravo
		}
		t.Members = append(t.Members, convertParticipant2TeamMember(p))
	}

	// TODO(haya14busa): Remove later when all participants data are in database.
	if len(m.TeamAlpha.Members) == 0 {
		fillInDummyMembers(false, m.TeamAlpha)
	}
	if len(m.TeamBravo.Members) == 0 {
		fillInDummyMembers(false, m.TeamBravo)
	}

	m.RoundName = roundName

	seenBattleOrders := make(map[int]bool)
	for _, b := range battles {
		m.Battles = append(m.Battles, convertBattle(b, m))
		seenBattleOrders[int(b.Order)] = true
	}

	// Fill in not-finished battles.
	maxBattleNum := getMaxBattleNum(match)
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
