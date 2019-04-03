// Code generated by go-swagger; DO NOT EDIT.

package gre_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateGreNetworkHandlerFunc turns a function with the right signature into a create gre network handler
type CreateGreNetworkHandlerFunc func(CreateGreNetworkParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateGreNetworkHandlerFunc) Handle(params CreateGreNetworkParams) middleware.Responder {
	return fn(params)
}

// CreateGreNetworkHandler interface for that can handle valid create gre network params
type CreateGreNetworkHandler interface {
	Handle(CreateGreNetworkParams) middleware.Responder
}

// NewCreateGreNetwork creates a new http.Handler for the create gre network operation
func NewCreateGreNetwork(ctx *middleware.Context, handler CreateGreNetworkHandler) *CreateGreNetwork {
	return &CreateGreNetwork{Context: ctx, Handler: handler}
}

/*CreateGreNetwork swagger:route POST /gre-networks greNetworks createGreNetwork

CreateGreNetwork create gre network API

*/
type CreateGreNetwork struct {
	Context *middleware.Context
	Handler CreateGreNetworkHandler
}

func (o *CreateGreNetwork) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateGreNetworkParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
