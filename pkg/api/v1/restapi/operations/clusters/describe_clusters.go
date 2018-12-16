// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DescribeClustersHandlerFunc turns a function with the right signature into a describe clusters handler
type DescribeClustersHandlerFunc func(DescribeClustersParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DescribeClustersHandlerFunc) Handle(params DescribeClustersParams) middleware.Responder {
	return fn(params)
}

// DescribeClustersHandler interface for that can handle valid describe clusters params
type DescribeClustersHandler interface {
	Handle(DescribeClustersParams) middleware.Responder
}

// NewDescribeClusters creates a new http.Handler for the describe clusters operation
func NewDescribeClusters(ctx *middleware.Context, handler DescribeClustersHandler) *DescribeClusters {
	return &DescribeClusters{Context: ctx, Handler: handler}
}

/*DescribeClusters swagger:route GET /clusters/describe clusters describeClusters

DescribeClusters describe clusters API

*/
type DescribeClusters struct {
	Context *middleware.Context
	Handler DescribeClustersHandler
}

func (o *DescribeClusters) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDescribeClustersParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
