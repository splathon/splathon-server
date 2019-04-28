// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"io"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/rs/cors"
	"github.com/splathon/splathon-server/splathon"
	"github.com/splathon/splathon-server/splathon/serror"
	"github.com/splathon/splathon-server/splathon/swagutils"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"github.com/splathon/splathon-server/swagger/restapi/operations/ranking"
	"github.com/splathon/splathon-server/swagger/restapi/operations/reception"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
)

//go:generate swagger generate server --target ../../swagger --name Splathon --spec ../../splathon-api/swagger.yaml

func configureFlags(api *operations.SplathonAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SplathonAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	thonHandler, err := splathon.NewDefaultHandler()
	if err != nil {
		log.Fatal(err)
	}

	api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
		res, err := thonHandler.Login(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return operations.NewLoginOK().WithPayload(res)
	})
	api.GetEventHandler = operations.GetEventHandlerFunc(func(params operations.GetEventParams) middleware.Responder {
		res, err := thonHandler.GetEvent(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return operations.NewGetEventOK().WithPayload(res)
	})
	api.MatchGetMatchHandler = match.GetMatchHandlerFunc(func(params match.GetMatchParams) middleware.Responder {
		res, err := thonHandler.GetMatch(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return match.NewGetMatchOK().WithPayload(res)
	})
	api.MatchGetNextMatchHandler = match.GetNextMatchHandlerFunc(func(params match.GetNextMatchParams) middleware.Responder {
		res, err := thonHandler.GetNextMatch(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return match.NewGetNextMatchOK().WithPayload(res)
	})
	api.ResultGetResultHandler = result.GetResultHandlerFunc(func(params result.GetResultParams) middleware.Responder {
		res, err := thonHandler.GetResult(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return result.NewGetResultOK().WithPayload(res)
	})
	api.RankingGetRankingHandler = ranking.GetRankingHandlerFunc(func(params ranking.GetRankingParams) middleware.Responder {
		res, err := thonHandler.GetRanking(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return ranking.NewGetRankingOK().WithPayload(res)
	})
	api.ListTeamsHandler = operations.ListTeamsHandlerFunc(func(params operations.ListTeamsParams) middleware.Responder {
		res, err := thonHandler.ListTeams(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return operations.NewListTeamsOK().WithPayload(res)
	})
	api.GetTeamDetailHandler = operations.GetTeamDetailHandlerFunc(func(params operations.GetTeamDetailParams) middleware.Responder {
		res, err := thonHandler.GetTeamDetail(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return operations.NewGetTeamDetailOK().WithPayload(res)
	})
	api.UpdateBattleHandler = operations.UpdateBattleHandlerFunc(func(params operations.UpdateBattleParams) middleware.Responder {
		if err := thonHandler.UpdateBattle(params.HTTPRequest.Context(), params); err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return operations.NewUpdateBattleOK()
	})
	api.GetParticipantsDataForReceptionHandler = operations.GetParticipantsDataForReceptionHandlerFunc(func(params operations.GetParticipantsDataForReceptionParams) middleware.Responder {
		res, err := thonHandler.GetParticipantsDataForReception(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return operations.NewGetParticipantsDataForReceptionOK().WithPayload(res)
	})
	api.CompleteReceptionHandler = operations.CompleteReceptionHandlerFunc(func(params operations.CompleteReceptionParams) middleware.Responder {
		if err := thonHandler.CompleteReception(params.HTTPRequest.Context(), params); err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return operations.NewCompleteReceptionOK()
	})
	api.ReceptionGetReceptionHandler = reception.GetReceptionHandlerFunc(func(params reception.GetReceptionParams) middleware.Responder {
		res, err := thonHandler.GetReception(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return reception.NewGetReceptionOK().WithPayload(res)
	})
	api.ListNoticesHandler = operations.ListNoticesHandlerFunc(func(params operations.ListNoticesParams) middleware.Responder {
		res, err := thonHandler.ListNotices(params.HTTPRequest.Context(), params)
		if err != nil {
			return swagutils.Error(err)
		}
		return operations.NewListNoticesOK().WithPayload(res)
	})
	api.AdminWriteNoticeHandler = admin.WriteNoticeHandlerFunc(func(params admin.WriteNoticeParams) middleware.Responder {
		if err := thonHandler.WriteNotice(params.HTTPRequest.Context(), params); err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return admin.NewWriteNoticeOK()
	})
	api.AdminDeleteNoticeHandler = admin.DeleteNoticeHandlerFunc(func(params admin.DeleteNoticeParams) middleware.Responder {
		return middleware.NotImplemented("operation admin.DeleteNotice has not yet been implemented")
	})
	api.AdminListReceptionHandler = admin.ListReceptionHandlerFunc(func(params admin.ListReceptionParams) middleware.Responder {
		res, err := thonHandler.ListReception(params.HTTPRequest.Context(), params)
		if err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return admin.NewListReceptionOK().WithPayload(res)
	})
	api.AdminUpdateReceptionHandler = admin.UpdateReceptionHandlerFunc(func(params admin.UpdateReceptionParams) middleware.Responder {
		if err := thonHandler.UpdateReception(params.HTTPRequest.Context(), params); err != nil {
			return logAndErr(err, params.HTTPRequest)
		}
		return admin.NewUpdateReceptionOK()
	})

	api.ServerShutdown = func() {
		if closer, ok := thonHandler.(io.Closer); ok {
			if err := closer.Close(); err != nil {
				api.Logger("failed to close handler: %v", err)
			}
		}
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func logAndErr(err error, req *http.Request) middleware.Responder {
	code := 500
	if splathonErr, ok := err.(*serror.Error); ok {
		code = splathonErr.Code
	}
	log.Printf("ERROR: code:%d\tmethod:%s\turl:%s\terror:%v", code, req.Method, req.URL.String(), err)
	return swagutils.Error(err)
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})
	return c.Handler(handler)
}
