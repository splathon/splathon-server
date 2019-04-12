// Code generated by go-swagger; DO NOT EDIT.

package result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetResultHandlerFunc turns a function with the right signature into a get result handler
type GetResultHandlerFunc func(GetResultParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetResultHandlerFunc) Handle(params GetResultParams) middleware.Responder {
	return fn(params)
}

// GetResultHandler interface for that can handle valid get result params
type GetResultHandler interface {
	Handle(GetResultParams) middleware.Responder
}

// NewGetResult creates a new http.Handler for the get result operation
func NewGetResult(ctx *middleware.Context, handler GetResultHandler) *GetResult {
	return &GetResult{Context: ctx, Handler: handler}
}

/*GetResult swagger:route GET /v{eventId}/results result getResult

リザルト一覧を返す。リザルトと言いつつ終了していない未来のマッチも返す。ゲスト・管理アプリ両方から使う。team_idを指定するとそのチームのみの結果が返ってくる。

*/
type GetResult struct {
	Context *middleware.Context
	Handler GetResultHandler
}

func (o *GetResult) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetResultParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
