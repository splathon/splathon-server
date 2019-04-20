// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateBattleHandlerFunc turns a function with the right signature into a update battle handler
type UpdateBattleHandlerFunc func(UpdateBattleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateBattleHandlerFunc) Handle(params UpdateBattleParams) middleware.Responder {
	return fn(params)
}

// UpdateBattleHandler interface for that can handle valid update battle params
type UpdateBattleHandler interface {
	Handle(UpdateBattleParams) middleware.Responder
}

// NewUpdateBattle creates a new http.Handler for the update battle operation
func NewUpdateBattle(ctx *middleware.Context, handler UpdateBattleHandler) *UpdateBattle {
	return &UpdateBattle{Context: ctx, Handler: handler}
}

/*UpdateBattle swagger:route POST /v{eventId}/matches/{matchId} match admin updateBattle

Update a battle data in the match.

*/
type UpdateBattle struct {
	Context *middleware.Context
	Handler UpdateBattleHandler
}

func (o *UpdateBattle) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateBattleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
