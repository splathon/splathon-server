package splathon

import (
	"context"

	"github.com/splathon/splathon-server/splathon/handler/pg"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"github.com/splathon/splathon-server/swagger/restapi/operations/ranking"
	"github.com/splathon/splathon-server/swagger/restapi/operations/reception"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
)

type Handler interface {
	GetEvent(context.Context, operations.GetEventParams) (*models.Event, error)
	GetResult(context.Context, result.GetResultParams) (*models.Results, error)
	GetMatch(context.Context, match.GetMatchParams) (*models.Match, error)
	UpdateMatch(context.Context, admin.UpdateMatchParams) error
	GetRanking(context.Context, ranking.GetRankingParams) (*models.Ranking, error)
	ListTeams(context.Context, operations.ListTeamsParams) (*models.Teams, error)
	GetSchedule(context.Context, operations.GetScheduleParams) (*models.Schedule, error)
	UpdateSchedule(context.Context, admin.UpdateScheduleParams) error
	GetTeamDetail(context.Context, operations.GetTeamDetailParams) (*models.Team, error)
	UpdateBattle(context.Context, operations.UpdateBattleParams) error
	Login(context.Context, operations.LoginParams) (*models.LoginResponse, error)
	GetReception(context.Context, reception.GetReceptionParams) (*models.ReceptionResponse, error)
	GetParticipantsDataForReception(context.Context, operations.GetParticipantsDataForReceptionParams) (*models.ReceptionPartcipantsDataResponse, error)
	CompleteReception(context.Context, operations.CompleteReceptionParams) error
	ListNotices(context.Context, operations.ListNoticesParams) (*models.ListNoticesResponse, error)
	WriteNotice(context.Context, admin.WriteNoticeParams) error
	DeleteNotice(context.Context, admin.DeleteNoticeParams) error
	GetNextMatch(context.Context, match.GetNextMatchParams) (*models.GetNextMatchResponse, error)
	ListReception(context.Context, admin.ListReceptionParams) (*models.ListReceptionResponse, error)
	UpdateReception(context.Context, admin.UpdateReceptionParams) error
	AddTournamentRound(context.Context, admin.AddTournamentRoundParams) error
	CreateNewQualifier(context.Context, admin.CreateNewQualifierParams) error
	DeleteQualifier(context.Context, admin.DeleteQualifierParams) error

	// Below API doesn't use postgres in pg handler.
	UpdateReleaseQualifier(context.Context, admin.UpdateReleaseQualifierParams) error
	GetReleaseQualifier(context.Context, admin.GetReleaseQualifierParams) (int32, error)
}

func NewDefaultHandler() (Handler, error) {
	opt, err := pg.NewOptionFromEnv()
	if err != nil {
		return nil, err
	}
	return pg.NewHandler(opt)
}
