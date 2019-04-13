package pg

import (
	"context"
	"errors"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/ranking"
)

func (h *Handler) GetRanking(ctx context.Context, params ranking.GetRankingParams) (*models.Ranking, error) {
	return nil, errors.New("not implemented")
}
