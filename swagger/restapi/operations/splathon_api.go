// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
	"github.com/splathon/splathon-server/swagger/restapi/operations/match"
	"github.com/splathon/splathon-server/swagger/restapi/operations/ranking"
	"github.com/splathon/splathon-server/swagger/restapi/operations/reception"
	"github.com/splathon/splathon-server/swagger/restapi/operations/result"
)

// NewSplathonAPI creates a new Splathon instance
func NewSplathonAPI(spec *loads.Document) *SplathonAPI {
	return &SplathonAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		AdminDeleteNoticeHandler: admin.DeleteNoticeHandlerFunc(func(params admin.DeleteNoticeParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminDeleteNotice has not yet been implemented")
		}),
		AdminUpdateReleaseQualifierHandler: admin.UpdateReleaseQualifierHandlerFunc(func(params admin.UpdateReleaseQualifierParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminUpdateReleaseQualifier has not yet been implemented")
		}),
		AdminAddTournamentRoundHandler: admin.AddTournamentRoundHandlerFunc(func(params admin.AddTournamentRoundParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminAddTournamentRound has not yet been implemented")
		}),
		CompleteReceptionHandler: CompleteReceptionHandlerFunc(func(params CompleteReceptionParams) middleware.Responder {
			return middleware.NotImplemented("operation CompleteReception has not yet been implemented")
		}),
		GetEventHandler: GetEventHandlerFunc(func(params GetEventParams) middleware.Responder {
			return middleware.NotImplemented("operation GetEvent has not yet been implemented")
		}),
		MatchGetMatchHandler: match.GetMatchHandlerFunc(func(params match.GetMatchParams) middleware.Responder {
			return middleware.NotImplemented("operation MatchGetMatch has not yet been implemented")
		}),
		MatchGetNextMatchHandler: match.GetNextMatchHandlerFunc(func(params match.GetNextMatchParams) middleware.Responder {
			return middleware.NotImplemented("operation MatchGetNextMatch has not yet been implemented")
		}),
		GetParticipantsDataForReceptionHandler: GetParticipantsDataForReceptionHandlerFunc(func(params GetParticipantsDataForReceptionParams) middleware.Responder {
			return middleware.NotImplemented("operation GetParticipantsDataForReception has not yet been implemented")
		}),
		RankingGetRankingHandler: ranking.GetRankingHandlerFunc(func(params ranking.GetRankingParams) middleware.Responder {
			return middleware.NotImplemented("operation RankingGetRanking has not yet been implemented")
		}),
		ReceptionGetReceptionHandler: reception.GetReceptionHandlerFunc(func(params reception.GetReceptionParams) middleware.Responder {
			return middleware.NotImplemented("operation ReceptionGetReception has not yet been implemented")
		}),
		AdminGetReleaseQualifierHandler: admin.GetReleaseQualifierHandlerFunc(func(params admin.GetReleaseQualifierParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminGetReleaseQualifier has not yet been implemented")
		}),
		ResultGetResultHandler: result.GetResultHandlerFunc(func(params result.GetResultParams) middleware.Responder {
			return middleware.NotImplemented("operation ResultGetResult has not yet been implemented")
		}),
		GetScheduleHandler: GetScheduleHandlerFunc(func(params GetScheduleParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSchedule has not yet been implemented")
		}),
		GetTeamDetailHandler: GetTeamDetailHandlerFunc(func(params GetTeamDetailParams) middleware.Responder {
			return middleware.NotImplemented("operation GetTeamDetail has not yet been implemented")
		}),
		ListNoticesHandler: ListNoticesHandlerFunc(func(params ListNoticesParams) middleware.Responder {
			return middleware.NotImplemented("operation ListNotices has not yet been implemented")
		}),
		AdminListReceptionHandler: admin.ListReceptionHandlerFunc(func(params admin.ListReceptionParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminListReception has not yet been implemented")
		}),
		ListTeamsHandler: ListTeamsHandlerFunc(func(params ListTeamsParams) middleware.Responder {
			return middleware.NotImplemented("operation ListTeams has not yet been implemented")
		}),
		LoginHandler: LoginHandlerFunc(func(params LoginParams) middleware.Responder {
			return middleware.NotImplemented("operation Login has not yet been implemented")
		}),
		UpdateBattleHandler: UpdateBattleHandlerFunc(func(params UpdateBattleParams) middleware.Responder {
			return middleware.NotImplemented("operation UpdateBattle has not yet been implemented")
		}),
		AdminUpdateMatchHandler: admin.UpdateMatchHandlerFunc(func(params admin.UpdateMatchParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminUpdateMatch has not yet been implemented")
		}),
		AdminUpdateReceptionHandler: admin.UpdateReceptionHandlerFunc(func(params admin.UpdateReceptionParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminUpdateReception has not yet been implemented")
		}),
		AdminWriteNoticeHandler: admin.WriteNoticeHandlerFunc(func(params admin.WriteNoticeParams) middleware.Responder {
			return middleware.NotImplemented("operation AdminWriteNotice has not yet been implemented")
		}),
	}
}

/*SplathonAPI Splathonで使うAPI */
type SplathonAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json; charset=utf-8" mime type
	JSONProducer runtime.Producer

	// AdminDeleteNoticeHandler sets the operation handler for the delete notice operation
	AdminDeleteNoticeHandler admin.DeleteNoticeHandler
	// AdminUpdateReleaseQualifierHandler sets the operation handler for the update release qualifier operation
	AdminUpdateReleaseQualifierHandler admin.UpdateReleaseQualifierHandler
	// AdminAddTournamentRoundHandler sets the operation handler for the add tournament round operation
	AdminAddTournamentRoundHandler admin.AddTournamentRoundHandler
	// CompleteReceptionHandler sets the operation handler for the complete reception operation
	CompleteReceptionHandler CompleteReceptionHandler
	// GetEventHandler sets the operation handler for the get event operation
	GetEventHandler GetEventHandler
	// MatchGetMatchHandler sets the operation handler for the get match operation
	MatchGetMatchHandler match.GetMatchHandler
	// MatchGetNextMatchHandler sets the operation handler for the get next match operation
	MatchGetNextMatchHandler match.GetNextMatchHandler
	// GetParticipantsDataForReceptionHandler sets the operation handler for the get participants data for reception operation
	GetParticipantsDataForReceptionHandler GetParticipantsDataForReceptionHandler
	// RankingGetRankingHandler sets the operation handler for the get ranking operation
	RankingGetRankingHandler ranking.GetRankingHandler
	// ReceptionGetReceptionHandler sets the operation handler for the get reception operation
	ReceptionGetReceptionHandler reception.GetReceptionHandler
	// AdminGetReleaseQualifierHandler sets the operation handler for the get release qualifier operation
	AdminGetReleaseQualifierHandler admin.GetReleaseQualifierHandler
	// ResultGetResultHandler sets the operation handler for the get result operation
	ResultGetResultHandler result.GetResultHandler
	// GetScheduleHandler sets the operation handler for the get schedule operation
	GetScheduleHandler GetScheduleHandler
	// GetTeamDetailHandler sets the operation handler for the get team detail operation
	GetTeamDetailHandler GetTeamDetailHandler
	// ListNoticesHandler sets the operation handler for the list notices operation
	ListNoticesHandler ListNoticesHandler
	// AdminListReceptionHandler sets the operation handler for the list reception operation
	AdminListReceptionHandler admin.ListReceptionHandler
	// ListTeamsHandler sets the operation handler for the list teams operation
	ListTeamsHandler ListTeamsHandler
	// LoginHandler sets the operation handler for the login operation
	LoginHandler LoginHandler
	// UpdateBattleHandler sets the operation handler for the update battle operation
	UpdateBattleHandler UpdateBattleHandler
	// AdminUpdateMatchHandler sets the operation handler for the update match operation
	AdminUpdateMatchHandler admin.UpdateMatchHandler
	// AdminUpdateReceptionHandler sets the operation handler for the update reception operation
	AdminUpdateReceptionHandler admin.UpdateReceptionHandler
	// AdminWriteNoticeHandler sets the operation handler for the write notice operation
	AdminWriteNoticeHandler admin.WriteNoticeHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *SplathonAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *SplathonAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *SplathonAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *SplathonAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *SplathonAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *SplathonAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *SplathonAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the SplathonAPI
func (o *SplathonAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.AdminDeleteNoticeHandler == nil {
		unregistered = append(unregistered, "admin.DeleteNoticeHandler")
	}

	if o.AdminUpdateReleaseQualifierHandler == nil {
		unregistered = append(unregistered, "admin.UpdateReleaseQualifierHandler")
	}

	if o.AdminAddTournamentRoundHandler == nil {
		unregistered = append(unregistered, "admin.AddTournamentRoundHandler")
	}

	if o.CompleteReceptionHandler == nil {
		unregistered = append(unregistered, "CompleteReceptionHandler")
	}

	if o.GetEventHandler == nil {
		unregistered = append(unregistered, "GetEventHandler")
	}

	if o.MatchGetMatchHandler == nil {
		unregistered = append(unregistered, "match.GetMatchHandler")
	}

	if o.MatchGetNextMatchHandler == nil {
		unregistered = append(unregistered, "match.GetNextMatchHandler")
	}

	if o.GetParticipantsDataForReceptionHandler == nil {
		unregistered = append(unregistered, "GetParticipantsDataForReceptionHandler")
	}

	if o.RankingGetRankingHandler == nil {
		unregistered = append(unregistered, "ranking.GetRankingHandler")
	}

	if o.ReceptionGetReceptionHandler == nil {
		unregistered = append(unregistered, "reception.GetReceptionHandler")
	}

	if o.AdminGetReleaseQualifierHandler == nil {
		unregistered = append(unregistered, "admin.GetReleaseQualifierHandler")
	}

	if o.ResultGetResultHandler == nil {
		unregistered = append(unregistered, "result.GetResultHandler")
	}

	if o.GetScheduleHandler == nil {
		unregistered = append(unregistered, "GetScheduleHandler")
	}

	if o.GetTeamDetailHandler == nil {
		unregistered = append(unregistered, "GetTeamDetailHandler")
	}

	if o.ListNoticesHandler == nil {
		unregistered = append(unregistered, "ListNoticesHandler")
	}

	if o.AdminListReceptionHandler == nil {
		unregistered = append(unregistered, "admin.ListReceptionHandler")
	}

	if o.ListTeamsHandler == nil {
		unregistered = append(unregistered, "ListTeamsHandler")
	}

	if o.LoginHandler == nil {
		unregistered = append(unregistered, "LoginHandler")
	}

	if o.UpdateBattleHandler == nil {
		unregistered = append(unregistered, "UpdateBattleHandler")
	}

	if o.AdminUpdateMatchHandler == nil {
		unregistered = append(unregistered, "admin.UpdateMatchHandler")
	}

	if o.AdminUpdateReceptionHandler == nil {
		unregistered = append(unregistered, "admin.UpdateReceptionHandler")
	}

	if o.AdminWriteNoticeHandler == nil {
		unregistered = append(unregistered, "admin.WriteNoticeHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *SplathonAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *SplathonAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *SplathonAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *SplathonAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *SplathonAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *SplathonAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the splathon API
func (o *SplathonAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *SplathonAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/v{eventId}/notices/{noticeId}"] = admin.NewDeleteNotice(o.context, o.AdminDeleteNoticeHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v{eventId}/release-qualifier"] = admin.NewUpdateReleaseQualifier(o.context, o.AdminUpdateReleaseQualifierHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v{eventId}/tournament"] = admin.NewAddTournamentRound(o.context, o.AdminAddTournamentRoundHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v{eventId}/reception/{splathonReceptionCode}/complete"] = NewCompleteReception(o.context, o.CompleteReceptionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/event"] = NewGetEvent(o.context, o.GetEventHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/matches/{matchId}"] = match.NewGetMatch(o.context, o.MatchGetMatchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/next-match"] = match.NewGetNextMatch(o.context, o.MatchGetNextMatchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/reception/{splathonReceptionCode}"] = NewGetParticipantsDataForReception(o.context, o.GetParticipantsDataForReceptionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/ranking"] = ranking.NewGetRanking(o.context, o.RankingGetRankingHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/reception"] = reception.NewGetReception(o.context, o.ReceptionGetReceptionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/release-qualifier"] = admin.NewGetReleaseQualifier(o.context, o.AdminGetReleaseQualifierHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/results"] = result.NewGetResult(o.context, o.ResultGetResultHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/schedule"] = NewGetSchedule(o.context, o.GetScheduleHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/teams/{team_id}"] = NewGetTeamDetail(o.context, o.GetTeamDetailHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/notices"] = NewListNotices(o.context, o.ListNoticesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/list-reception"] = admin.NewListReception(o.context, o.AdminListReceptionHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/v{eventId}/teams"] = NewListTeams(o.context, o.ListTeamsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v{eventId}/login"] = NewLogin(o.context, o.LoginHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v{eventId}/matches/{matchId}"] = NewUpdateBattle(o.context, o.UpdateBattleHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/v{eventId}/matches/{matchId}"] = admin.NewUpdateMatch(o.context, o.AdminUpdateMatchHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v{eventId}/update-reception"] = admin.NewUpdateReception(o.context, o.AdminUpdateReceptionHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/v{eventId}/notices"] = admin.NewWriteNotice(o.context, o.AdminWriteNoticeHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *SplathonAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *SplathonAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *SplathonAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *SplathonAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
