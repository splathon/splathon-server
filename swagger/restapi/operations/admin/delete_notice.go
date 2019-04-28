// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteNoticeHandlerFunc turns a function with the right signature into a delete notice handler
type DeleteNoticeHandlerFunc func(DeleteNoticeParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteNoticeHandlerFunc) Handle(params DeleteNoticeParams) middleware.Responder {
	return fn(params)
}

// DeleteNoticeHandler interface for that can handle valid delete notice params
type DeleteNoticeHandler interface {
	Handle(DeleteNoticeParams) middleware.Responder
}

// NewDeleteNotice creates a new http.Handler for the delete notice operation
func NewDeleteNotice(ctx *middleware.Context, handler DeleteNoticeHandler) *DeleteNotice {
	return &DeleteNotice{Context: ctx, Handler: handler}
}

/*DeleteNotice swagger:route DELETE /v{eventId}/notices/{noticeId} admin deleteNotice

DeleteNotice delete notice API

*/
type DeleteNotice struct {
	Context *middleware.Context
	Handler DeleteNoticeHandler
}

func (o *DeleteNotice) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteNoticeParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
