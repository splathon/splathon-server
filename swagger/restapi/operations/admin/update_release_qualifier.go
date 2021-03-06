// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UpdateReleaseQualifierHandlerFunc turns a function with the right signature into a update release qualifier handler
type UpdateReleaseQualifierHandlerFunc func(UpdateReleaseQualifierParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateReleaseQualifierHandlerFunc) Handle(params UpdateReleaseQualifierParams) middleware.Responder {
	return fn(params)
}

// UpdateReleaseQualifierHandler interface for that can handle valid update release qualifier params
type UpdateReleaseQualifierHandler interface {
	Handle(UpdateReleaseQualifierParams) middleware.Responder
}

// NewUpdateReleaseQualifier creates a new http.Handler for the update release qualifier operation
func NewUpdateReleaseQualifier(ctx *middleware.Context, handler UpdateReleaseQualifierHandler) *UpdateReleaseQualifier {
	return &UpdateReleaseQualifier{Context: ctx, Handler: handler}
}

/*UpdateReleaseQualifier swagger:route PUT /v{eventId}/release-qualifier admin updateReleaseQualifier

UpdateReleaseQualifier update release qualifier API

*/
type UpdateReleaseQualifier struct {
	Context *middleware.Context
	Handler UpdateReleaseQualifierHandler
}

func (o *UpdateReleaseQualifier) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateReleaseQualifierParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
