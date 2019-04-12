// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/splathon/splathon-server/splathon"
	"github.com/splathon/splathon-server/splathon/swagutils"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
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

	api.MatchGetMatchHandler = match.GetMatchHandlerFunc(func(params match.GetMatchParams) middleware.Responder {
		res, err := thonHandler.GetMatch(context.TODO(), params)
		if err != nil {
			return swagutils.Error(err)
		}
		return match.NewGetMatchOK().WithPayload(res)
	})
	api.ResultGetResultHandler = result.GetResultHandlerFunc(func(params result.GetResultParams) middleware.Responder {
		res, err := thonHandler.GetResult(context.TODO(), params)
		if err != nil {
			return swagutils.Error(err)
		}
		return result.NewGetResultOK().WithPayload(res)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
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
	return handler
}
