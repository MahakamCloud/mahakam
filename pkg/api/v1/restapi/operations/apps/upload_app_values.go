// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UploadAppValuesHandlerFunc turns a function with the right signature into a upload app values handler
type UploadAppValuesHandlerFunc func(UploadAppValuesParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UploadAppValuesHandlerFunc) Handle(params UploadAppValuesParams) middleware.Responder {
	return fn(params)
}

// UploadAppValuesHandler interface for that can handle valid upload app values params
type UploadAppValuesHandler interface {
	Handle(UploadAppValuesParams) middleware.Responder
}

// NewUploadAppValues creates a new http.Handler for the upload app values operation
func NewUploadAppValues(ctx *middleware.Context, handler UploadAppValuesHandler) *UploadAppValues {
	return &UploadAppValues{Context: ctx, Handler: handler}
}

/*UploadAppValues swagger:route POST /apps/values apps uploadAppValues

UploadAppValues upload app values API

*/
type UploadAppValues struct {
	Context *middleware.Context
	Handler UploadAppValuesHandler
}

func (o *UploadAppValues) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUploadAppValuesParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
