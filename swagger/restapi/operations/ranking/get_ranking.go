// Code generated by go-swagger; DO NOT EDIT.

package ranking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRankingHandlerFunc turns a function with the right signature into a get ranking handler
type GetRankingHandlerFunc func(GetRankingParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRankingHandlerFunc) Handle(params GetRankingParams) middleware.Responder {
	return fn(params)
}

// GetRankingHandler interface for that can handle valid get ranking params
type GetRankingHandler interface {
	Handle(GetRankingParams) middleware.Responder
}

// NewGetRanking creates a new http.Handler for the get ranking operation
func NewGetRanking(ctx *middleware.Context, handler GetRankingHandler) *GetRanking {
	return &GetRanking{Context: ctx, Handler: handler}
}

/*GetRanking swagger:route GET /v{eventId}/ranking ranking getRanking

予選ランキングを返す。

*/
type GetRanking struct {
	Context *middleware.Context
	Handler GetRankingHandler
}

func (o *GetRanking) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRankingParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
