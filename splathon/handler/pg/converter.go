package pg

import (
	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/splathon/spldata"
	"github.com/splathon/splathon-server/swagger/models"
)

func convertTeam(t *Team) *models.Team {
	return &models.Team{
		ID:   swag.Int32(int32(t.Id)),
		Name: swag.String(t.Name),
	}
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

func convertBattle(b *Battle, m *models.Match) *models.Battle {
	result := &models.Battle{
		ID:    swag.Int32(int32(b.Id)),
		Order: swag.Int32(b.Order),
	}
	switch b.WinnerId {
	case int64(*m.TeamAlpha.ID):
		result.Winner = "alpha"
	case int64(*m.TeamBravo.ID):
		result.Winner = "bravo"
	}
	if b.StageId.Valid {
		if n, ok := spldata.GetStageByID(int(b.StageId.Int64)); ok {
			result.Stage = &models.BattleStage{ID: int32(b.StageId.Int64), Name: n}
		}
	}
	if b.RuleId.Valid {
		if key, name, ok := spldata.GetRuleByID(int(b.RuleId.Int64)); ok {
			result.Rule = &models.BattleRule{Key: key, Name: name}
		}
	}
	return result
}
