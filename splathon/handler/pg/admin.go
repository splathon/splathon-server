package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/splathon/splathon-server/splathon/spldata"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"golang.org/x/sync/errgroup"
)

func (h *Handler) UpdateBattle(ctx context.Context, params operations.UpdateBattleParams) error {
	// TODO(haya14busa): implement LOGIN API.
	if os.Getenv("SPLATHON_APP_READONLY") == "1" {
		return errors.New("SPLATHON_APP_READONLY=1: write admin API is not available")
	}

	var eg errgroup.Group
	var (
		match Match
	)
	eg.Go(func() error {
		return h.db.Where("id = ?", params.MatchID).Find(&match).Error
	})
	if err := eg.Wait(); err != nil {
		return err
	}

	if params.Battle.Rule == nil {
		return errors.New("rule key is required")
	}
	rule, ok := spldata.GetRuleByKey(*params.Battle.Rule.Key)
	if !ok {
		return fmt.Errorf("invalid rule key: %q", *params.Battle.Rule.Key)
	}

	if params.Battle.Stage == nil {
		return errors.New("stage id is required")
	}

	battle := Battle{
		Order:   *params.Battle.Order,
		RuleId:  int64(rule.ID),
		StageId: int64(*params.Battle.Stage.ID),
	}
	switch params.Battle.Winner {
	case models.BattleWinnerAlpha:
		battle.WinnerId = sql.NullInt64{Int64: match.TeamId, Valid: true}
	case models.BattleWinnerBravo:
		battle.WinnerId = sql.NullInt64{Int64: match.OpponentId, Valid: true}
	}
	var res Battle
	query := Battle{
		Order:   *params.Battle.Order,
		MatchId: match.Id,
	}
	return h.db.Where(query).Assign(&battle).FirstOrCreate(&res).Error
}
