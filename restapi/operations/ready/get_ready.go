// Code generated by go-swagger; DO NOT EDIT.

package ready

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetReadyHandlerFunc turns a function with the right signature into a get ready handler
type GetReadyHandlerFunc func(GetReadyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReadyHandlerFunc) Handle(params GetReadyParams) middleware.Responder {
	return fn(params)
}

// GetReadyHandler interface for that can handle valid get ready params
type GetReadyHandler interface {
	Handle(GetReadyParams) middleware.Responder
}

// NewGetReady creates a new http.Handler for the get ready operation
func NewGetReady(ctx *middleware.Context, handler GetReadyHandler) *GetReady {
	return &GetReady{Context: ctx, Handler: handler}
}

/* GetReady swagger:route GET /ready ready getReady

verify that the API is running and ready to accept requests

*/
type GetReady struct {
	Context *middleware.Context
	Handler GetReadyHandler
}

func (o *GetReady) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReadyParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
