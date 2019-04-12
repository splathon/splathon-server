package pg

import (
	"context"
	"errors"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
)

// Handler is splathon API handler backed by PostgreSQL.
type Handler struct{}

func NewHandler() (*Handler, error) {
	return &Handler{}, nil
}

func (h *Handler) GetResult(ctx context.Context, params result.GetResultParams) (*models.Results, error) {
	return nil, errors.New("not implemented")
}

func (h *Handler) GetMatch(ctx context.Context, params match.GetMatchParams) (*models.Match, error) {
	return nil, errors.New("not implemented")
}
