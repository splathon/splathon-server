package pg

import (
	"context"
	"errors"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) Login(ctx context.Context, params operations.LoginParams) (*models.LoginResponse, error) {
	return nil, errors.New("operation .Login has not yet been implemented")
}
