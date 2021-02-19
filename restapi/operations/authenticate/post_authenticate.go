// Code generated by go-swagger; DO NOT EDIT.

package authenticate

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostAuthenticateHandlerFunc turns a function with the right signature into a post authenticate handler
type PostAuthenticateHandlerFunc func(PostAuthenticateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostAuthenticateHandlerFunc) Handle(params PostAuthenticateParams) middleware.Responder {
	return fn(params)
}

// PostAuthenticateHandler interface for that can handle valid post authenticate params
type PostAuthenticateHandler interface {
	Handle(PostAuthenticateParams) middleware.Responder
}

// NewPostAuthenticate creates a new http.Handler for the post authenticate operation
func NewPostAuthenticate(ctx *middleware.Context, handler PostAuthenticateHandler) *PostAuthenticate {
	return &PostAuthenticate{Context: ctx, Handler: handler}
}

/* PostAuthenticate swagger:route POST /authenticate authenticate postAuthenticate

authenticate user

authenticate user

*/
type PostAuthenticate struct {
	Context *middleware.Context
	Handler PostAuthenticateHandler
}

func (o *PostAuthenticate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostAuthenticateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
