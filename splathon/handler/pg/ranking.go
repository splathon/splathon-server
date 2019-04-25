package pg

import (
	"context"
	"fmt"
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
		teams        []*Team
		matches      []*Match
		participants []*Participant
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

	eg.Go(func() error {
		return h.db.Where("event_id = ? AND team_id IS NOT NULL", eventID).Order("id asc").Find(&participants).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	rankResp, err := buildRanking(teams, filterCompletedMatches(teams, matches), buildTeam2Members(participants)), nil
	if err != nil {
		return nil, err
	}
	if len(rankResp.Rankings) > 0 {
		rankResp.RankTime = fmt.Sprintf("予選第%dラウンド終了時点", rankResp.Rankings[0].NumOfMatches)
	} else {
		rankResp.RankTime = "開始時点"
	}
	return rankResp, nil
}

type teamResult struct {
	teamID          int64
	opponentTeamIDs []int64
	totalPoint      int64
}

func filterCompletedMatches(teams []*Team, matches []*Match) []*Match {
	ms := make([]*Match, 0, len(matches))
	completedQIDs := completedQualifierIDs(teams, matches)
	for _, m := range matches {
		if completedQIDs[m.QualifierId.Int64] {
			ms = append(ms, m)
		}
	}
	return ms
}

func completedQualifierIDs(teams []*Team, matches []*Match) map[int64]bool {
	q2tc := make(map[int64]int) // QualifierId to Team Counts
	for _, m := range matches {
		if m.TeamPoints == 0 && m.OpponentPoints == 0 {
			// Skip matches which has not been done yet.
			continue
		}
		q2tc[m.QualifierId.Int64] += 2
	}
	teamNum := len(teams)
	qids := make(map[int64]bool)
	for q, tc := range q2tc {
		if tc == teamNum {
			qids[q] = true
		}
	}
	return qids
}

func buildRanking(teams []*Team, matches []*Match, team2members map[int64][]*models.Member) *models.Ranking {
	teamMap := make(map[int64]*teamResult)
	for _, t := range teams {
		teamMap[t.Id] = &teamResult{
			teamID:          t.Id,
			opponentTeamIDs: make([]int64, 0),
		}
	}
	for _, m := range matches {
		teamMap[m.TeamId].opponentTeamIDs = append(teamMap[m.TeamId].opponentTeamIDs, m.OpponentId)
		teamMap[m.TeamId].totalPoint += m.TeamPoints
		teamMap[m.OpponentId].opponentTeamIDs = append(teamMap[m.OpponentId].opponentTeamIDs, m.TeamId)
		teamMap[m.OpponentId].totalPoint += m.OpponentPoints
	}

	rs := make([]*models.Rank, 0, len(teams))
	for _, t := range teams {
		team := convertTeam(t)

		if ms, ok := team2members[t.Id]; ok {
			team.Members = ms
		} else {
			// TODO(haya14busa): Remove later when all participants data are in database.
			fillInDummyMembers(false, team)
		}

		rank := &models.Rank{
			Team:         team,
			Point:        swag.Int32(int32(teamMap[t.Id].totalPoint)),
			Omwp:         omwp(t.Id, teamMap),
			NumOfMatches: int32(len(teamMap[t.Id].opponentTeamIDs)),
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
func omwp(teamID int64, teamMap map[int64]*teamResult) float64 {
	if len(teamMap[teamID].opponentTeamIDs) == 0 {
		return 0
	}
	sum := 0.0
	for _, opID := range teamMap[teamID].opponentTeamIDs {
		sum += float64(teamMap[opID].totalPoint) / float64(len(teamMap[opID].opponentTeamIDs)*3)
	}
	return sum / float64(len(teamMap[teamID].opponentTeamIDs))
}

const eps = 0.00000001

func floatEquals(a, b float64) bool {
	return math.Abs(a-b) < eps
}
