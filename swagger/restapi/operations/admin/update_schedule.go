// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateScheduleHandlerFunc turns a function with the right signature into a update schedule handler
type UpdateScheduleHandlerFunc func(UpdateScheduleParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateScheduleHandlerFunc) Handle(params UpdateScheduleParams) middleware.Responder {
	return fn(params)
}

// UpdateScheduleHandler interface for that can handle valid update schedule params
type UpdateScheduleHandler interface {
	Handle(UpdateScheduleParams) middleware.Responder
}

// NewUpdateSchedule creates a new http.Handler for the update schedule operation
func NewUpdateSchedule(ctx *middleware.Context, handler UpdateScheduleHandler) *UpdateSchedule {
	return &UpdateSchedule{Context: ctx, Handler: handler}
}

/*UpdateSchedule swagger:route PUT /v{eventId}/schedule admin updateSchedule

Update event schedule data

*/
type UpdateSchedule struct {
	Context *middleware.Context
	Handler UpdateScheduleHandler
}

func (o *UpdateSchedule) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateScheduleParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
