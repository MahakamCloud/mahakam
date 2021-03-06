// Code generated by go-swagger; DO NOT EDIT.

package gre_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetGreNetworksHandlerFunc turns a function with the right signature into a get gre networks handler
type GetGreNetworksHandlerFunc func(GetGreNetworksParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetGreNetworksHandlerFunc) Handle(params GetGreNetworksParams) middleware.Responder {
	return fn(params)
}

// GetGreNetworksHandler interface for that can handle valid get gre networks params
type GetGreNetworksHandler interface {
	Handle(GetGreNetworksParams) middleware.Responder
}

// NewGetGreNetworks creates a new http.Handler for the get gre networks operation
func NewGetGreNetworks(ctx *middleware.Context, handler GetGreNetworksHandler) *GetGreNetworks {
	return &GetGreNetworks{Context: ctx, Handler: handler}
}

/*GetGreNetworks swagger:route GET /gre-networks gre-networks getGreNetworks

GetGreNetworks get gre networks API

*/
type GetGreNetworks struct {
	Context *middleware.Context
	Handler GetGreNetworksHandler
}

func (o *GetGreNetworks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetGreNetworksParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
