package pg

import (
	"context"
	"math"
	"sort"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/ranking"
	"golang.org/x/sync/errgroup"
)

func (h *Handler) GetRanking(ctx context.Context, params ranking.GetRankingParams) (*models.Ranking, error) {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	var (
		teams   []*Team
		matches []*Match
	)

	var eg errgroup.Group

	eg.Go(func() error {
		qids := h.db.Table("qualifiers").Select("id").Where("event_id = ?", eventID).QueryExpr()
		query := "qualifier_id in (?) AND team_points IS NOT NULL AND opponent_points IS NOT NULL"
		return h.db.Where(query, qids).Find(&matches).Error
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Find(&teams).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return buildRanking(teams, matches), nil
}

type teamResult struct {
	teamID          int64
	opponentTeamIDs []int64
	totalPoint      float64
}

func buildRanking(teams []*Team, matches []*Match) *models.Ranking {
	teamMap := make(map[int64]*teamResult)
	for _, t := range teams {
		teamMap[t.Id] = &teamResult{
			teamID:          t.Id,
			opponentTeamIDs: make([]int64, 0),
			totalPoint:      float64(t.Points),
		}
	}
	for _, m := range matches {
		if m.TeamPoints == 0 && m.OpponentPoints == 0 {
			// Skip matches which has not been done yet.
			continue
		}
		teamMap[m.TeamId].opponentTeamIDs = append(teamMap[m.TeamId].opponentTeamIDs, m.OpponentId)
		teamMap[m.OpponentId].opponentTeamIDs = append(teamMap[m.OpponentId].opponentTeamIDs, m.TeamId)
	}

	rs := make([]*models.Rank, 0, len(teams))
	for _, t := range teams {
		rank := &models.Rank{
			Team:  convertTeam(t),
			Point: swag.Int32(t.Points),
			Omwp:  omwp(t.Id, teamMap),
		}
		rs = append(rs, rank)
	}
	sort.SliceStable(rs, func(i, j int) bool {
		ip := *rs[i].Point
		jp := *rs[j].Point
		if ip == jp {
			return rs[i].Omwp > rs[j].Omwp
		}
		return ip > jp
	})
	for i, r := range rs {
		if i == 0 {
			r.Rank = swag.Int32(1)
			continue
		}
		if *rs[i-1].Point == *rs[i].Point && floatEquals(rs[i-1].Omwp, rs[i].Omwp) {
			r.Rank = rs[i-1].Rank
			continue
		}
		r.Rank = swag.Int32(int32(i + 1))
	}
	return &models.Ranking{Rankings: rs}
}

// ref: https://dic.nicovideo.jp/a/%E3%82%B9%E3%82%A4%E3%82%B9%E3%83%89%E3%83%AD%E3%83%BC
func omwp(teamId int64, teamMap map[int64]*teamResult) float64 {
	var sum float64 = 0
	for _, opID := range teamMap[teamId].opponentTeamIDs {
		sum += teamMap[opID].totalPoint / float64(len(teamMap[opID].opponentTeamIDs)*3)
	}
	return sum / float64(len(teamMap[teamId].opponentTeamIDs))
}

const eps = 0.00000001

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < eps
}
