// Code generated by go-swagger; DO NOT EDIT.

package reception

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetVEventIDReceptionHandlerFunc turns a function with the right signature into a get v event ID reception handler
type GetVEventIDReceptionHandlerFunc func(GetVEventIDReceptionParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetVEventIDReceptionHandlerFunc) Handle(params GetVEventIDReceptionParams) middleware.Responder {
	return fn(params)
}

// GetVEventIDReceptionHandler interface for that can handle valid get v event ID reception params
type GetVEventIDReceptionHandler interface {
	Handle(GetVEventIDReceptionParams) middleware.Responder
}

// NewGetVEventIDReception creates a new http.Handler for the get v event ID reception operation
func NewGetVEventIDReception(ctx *middleware.Context, handler GetVEventIDReceptionHandler) *GetVEventIDReception {
	return &GetVEventIDReception{Context: ctx, Handler: handler}
}

/*GetVEventIDReception swagger:route GET /v{eventId}/reception reception getVEventIdReception

GetVEventIDReception get v event ID reception API

*/
type GetVEventIDReception struct {
	Context *middleware.Context
	Handler GetVEventIDReceptionHandler
}

func (o *GetVEventIDReception) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetVEventIDReceptionParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
