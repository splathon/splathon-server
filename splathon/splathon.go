package splathon

import (
	"context"

	"github.com/splathon/splathon-server/splathon/handler/pg"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"github.com/splathon/splathon-server/swagger/restapi/operations/ranking"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
)

type Handler interface {
	GetEvent(context.Context, operations.GetEventParams) (*models.Event, error)
	GetResult(context.Context, result.GetResultParams) (*models.Results, error)
	GetMatch(context.Context, match.GetMatchParams) (*models.Match, error)
	GetRanking(context.Context, ranking.GetRankingParams) (*models.Ranking, error)
	ListTeams(context.Context, operations.ListTeamsParams) (*models.Teams, error)
	UpdateBattle(context.Context, operations.UpdateBattleParams) error
	Login(context.Context, operations.LoginParams) (*models.LoginResponse, error)
}

func NewDefaultHandler() (Handler, error) {
	opt, err := pg.NewOptionFromEnv()
	if err != nil {
		return nil, err
	}
	return pg.NewHandler(opt)
}
