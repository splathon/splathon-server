package pg

import (
	"context"
	"errors"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
)

func (h *Handler) GetMatch(ctx context.Context, params match.GetMatchParams) (*models.Match, error) {
	return nil, errors.New("not implemented")
}
