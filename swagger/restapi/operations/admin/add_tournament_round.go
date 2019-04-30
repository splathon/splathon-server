// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AddTournamentRoundHandlerFunc turns a function with the right signature into a add tournament round handler
type AddTournamentRoundHandlerFunc func(AddTournamentRoundParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTournamentRoundHandlerFunc) Handle(params AddTournamentRoundParams) middleware.Responder {
	return fn(params)
}

// AddTournamentRoundHandler interface for that can handle valid add tournament round params
type AddTournamentRoundHandler interface {
	Handle(AddTournamentRoundParams) middleware.Responder
}

// NewAddTournamentRound creates a new http.Handler for the add tournament round operation
func NewAddTournamentRound(ctx *middleware.Context, handler AddTournamentRoundHandler) *AddTournamentRound {
	return &AddTournamentRound{Context: ctx, Handler: handler}
}

/*AddTournamentRound swagger:route POST /v{eventId}/tournament/ admin addTournamentRound

AddTournamentRound add tournament round API

*/
type AddTournamentRound struct {
	Context *middleware.Context
	Handler AddTournamentRoundHandler
}

func (o *AddTournamentRound) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddTournamentRoundParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}