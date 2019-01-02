// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AllocateOrReleaseFromIPPoolHandlerFunc turns a function with the right signature into a allocate or release from Ip pool handler
type AllocateOrReleaseFromIPPoolHandlerFunc func(AllocateOrReleaseFromIPPoolParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AllocateOrReleaseFromIPPoolHandlerFunc) Handle(params AllocateOrReleaseFromIPPoolParams) middleware.Responder {
	return fn(params)
}

// AllocateOrReleaseFromIPPoolHandler interface for that can handle valid allocate or release from Ip pool params
type AllocateOrReleaseFromIPPoolHandler interface {
	Handle(AllocateOrReleaseFromIPPoolParams) middleware.Responder
}

// NewAllocateOrReleaseFromIPPool creates a new http.Handler for the allocate or release from Ip pool operation
func NewAllocateOrReleaseFromIPPool(ctx *middleware.Context, handler AllocateOrReleaseFromIPPoolHandler) *AllocateOrReleaseFromIPPool {
	return &AllocateOrReleaseFromIPPool{Context: ctx, Handler: handler}
}

/*AllocateOrReleaseFromIPPool swagger:route POST /networks/pools/ipPools/{poolId} networks allocateOrReleaseFromIpPool

AllocateOrReleaseFromIPPool allocate or release from Ip pool API

*/
type AllocateOrReleaseFromIPPool struct {
	Context *middleware.Context
	Handler AllocateOrReleaseFromIPPoolHandler
}

func (o *AllocateOrReleaseFromIPPool) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAllocateOrReleaseFromIPPoolParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}