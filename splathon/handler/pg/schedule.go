package pg

import (
	"context"
	"fmt"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
)

func (h *Handler) GetSchedule(ctx context.Context, params operations.GetScheduleParams) (*models.Schedule, error) {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}
	return GetEventSchedule(ctx, eventID)
}

func (h *Handler) UpdateSchedule(ctx context.Context, params admin.UpdateScheduleParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return err
	}
	if err := UpdateEventSchedule(ctx, eventID, params.Request); err != nil {
		return fmt.Errorf("failed to update schedule: %v", err)
	}
	return nil
}
