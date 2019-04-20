package pg

import (
	"context"
	"errors"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) GetEvent(ctx context.Context, params operations.GetEventParams) (*models.Event, error) {
	return nil, errors.New("GetEvent is not implemented yet")
}
