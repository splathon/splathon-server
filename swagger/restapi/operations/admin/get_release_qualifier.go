// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReleaseQualifierHandlerFunc turns a function with the right signature into a get release qualifier handler
type GetReleaseQualifierHandlerFunc func(GetReleaseQualifierParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReleaseQualifierHandlerFunc) Handle(params GetReleaseQualifierParams) middleware.Responder {
	return fn(params)
}

// GetReleaseQualifierHandler interface for that can handle valid get release qualifier params
type GetReleaseQualifierHandler interface {
	Handle(GetReleaseQualifierParams) middleware.Responder
}

// NewGetReleaseQualifier creates a new http.Handler for the get release qualifier operation
func NewGetReleaseQualifier(ctx *middleware.Context, handler GetReleaseQualifierHandler) *GetReleaseQualifier {
	return &GetReleaseQualifier{Context: ctx, Handler: handler}
}

/*GetReleaseQualifier swagger:route GET /v{eventId}/release-qualifier admin getReleaseQualifier

GetReleaseQualifier get release qualifier API

*/
type GetReleaseQualifier struct {
	Context *middleware.Context
	Handler GetReleaseQualifierHandler
}

func (o *GetReleaseQualifier) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReleaseQualifierParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}