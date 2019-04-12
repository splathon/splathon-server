package splathon

import (
	"context"

	"github.com/splathon/splathon-server/splathon/handler/pg"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
)

type Handler interface {
	GetResult(context.Context, result.GetResultParams) (*models.Results, error)
	GetMatch(context.Context, match.GetMatchParams) (*models.Match, error)
}

func NewDefaultHandler() (Handler, error) {
	return pg.NewHandler()
}
