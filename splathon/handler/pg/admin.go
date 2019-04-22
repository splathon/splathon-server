package pg

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/splathon/splathon-server/splathon/spldata"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"golang.org/x/sync/errgroup"
)

const (
	qualifierWinPt  = 3
	qualifierLosePt = 0
	qualifierDrawPt = 1
)

func (h *Handler) UpdateBattle(ctx context.Context, params operations.UpdateBattleParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
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

	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := h.updateBattleAndMatch(ctx, tx, updateBattleAndMatchParams{
		updateBattleParams: params,
		matchID:            match.Id,
		order:              *params.Battle.Order,
		ruleID:             rule.ID,
		stageID:            int(*params.Battle.Stage.ID),
		alphaTeamID:        match.TeamId,
		bravoTeamID:        match.OpponentId,
	}, match); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

type updateBattleAndMatchParams struct {
	updateBattleParams operations.UpdateBattleParams
	matchID            int64
	order              int32
	ruleID             int
	stageID            int
	alphaTeamID        int64
	bravoTeamID        int64
}

func (h *Handler) updateBattleAndMatch(ctx context.Context, tx *gorm.DB, params updateBattleAndMatchParams, match Match) error {
	// Update the battle result.
	battle := Battle{
		Order:   params.order,
		RuleId:  sql.NullInt64{Int64: int64(params.ruleID), Valid: true},
		StageId: sql.NullInt64{Int64: int64(params.stageID), Valid: true},
	}
	switch params.updateBattleParams.Battle.Winner {
	case models.BattleWinnerAlpha:
		battle.WinnerId = sql.NullInt64{Int64: params.alphaTeamID, Valid: true}
	case models.BattleWinnerBravo:
		battle.WinnerId = sql.NullInt64{Int64: params.bravoTeamID, Valid: true}
	}
	var res Battle
	query := Battle{
		Order:   params.order,
		MatchId: params.matchID,
	}
	if err := tx.Where(query).Assign(&battle).FirstOrCreate(&res).Error; err != nil {
		return err
	}

	// Fetch battles.
	var battles []*Battle
	if err := tx.Where("match_id = ?", params.matchID).Find(&battles).Error; err != nil {
		return err
	}

	maxBattleNum := getMaxBattleNum(match)
	alphaWin := 0
	bravoWin := 0
	for _, b := range battles {
		if !b.WinnerId.Valid {
			continue
		}
		switch b.WinnerId.Int64 {
		case match.TeamId:
			alphaWin++
		case match.OpponentId:
			bravoWin++
		}
	}

	if match.QualifierId.Valid && alphaWin+bravoWin == maxBattleNum {
		// Update the match in qualifier.
		if alphaWin > bravoWin {
			match.TeamPoints = qualifierWinPt
			match.OpponentPoints = qualifierLosePt
		} else if alphaWin < bravoWin {
			match.TeamPoints = qualifierLosePt
			match.OpponentPoints = qualifierWinPt
		} else {
			match.TeamPoints = qualifierDrawPt
			match.OpponentPoints = qualifierDrawPt
		}
		tx.Save(match)
	}
	return nil
}
