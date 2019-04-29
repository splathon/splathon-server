package pg

import (
	"context"
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
)

func (h *Handler) AddTournamentRound(ctx context.Context, params admin.AddTournamentRoundParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}
	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := h.addTournamentRoundTx(ctx, tx, params); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (h *Handler) addTournamentRoundTx(ctx context.Context, tx *gorm.DB, params admin.AddTournamentRoundParams) error {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return err
	}
	tournament := &Tournament{
		EventId:   eventID,
		Round:     *params.Request.Round,
		Name:      *params.Request.RoundName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if err := tx.Create(&tournament).Error; err != nil {
		return err
	}
	for _, m := range params.Request.Matches {
		newMatch := &Match{
			TournamentId: sql.NullInt64{Int64: tournament.Id, Valid: true},
			TeamId:       *m.AlphaTeamID,
			OpponentId:   *m.BravoTeamID,
			RoomId:       *m.RoomID,
			Order:        int64(*m.OrderInRoom),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}
		if err := tx.Create(&newMatch).Error; err != nil {
			return err
		}
	}
	return nil
}
