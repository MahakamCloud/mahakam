// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/mahakamcloud/mahakam/pkg/api/v1/models"
)

// GetNodesOKCode is the HTTP code returned for type GetNodesOK
const GetNodesOKCode int = 200

/*GetNodesOK list created nodes

swagger:response getNodesOK
*/
type GetNodesOK struct {

	/*
	  In: Body
	*/
	Payload []*models.Node `json:"body,omitempty"`
}

// NewGetNodesOK creates GetNodesOK with default headers values
func NewGetNodesOK() *GetNodesOK {

	return &GetNodesOK{}
}

// WithPayload adds the payload to the get nodes o k response
func (o *GetNodesOK) WithPayload(payload []*models.Node) *GetNodesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nodes o k response
func (o *GetNodesOK) SetPayload(payload []*models.Node) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*models.Node, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*GetNodesDefault generic error response

swagger:response getNodesDefault
*/
type GetNodesDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetNodesDefault creates GetNodesDefault with default headers values
func NewGetNodesDefault(code int) *GetNodesDefault {
	if code <= 0 {
		code = 500
	}

	return &GetNodesDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get nodes default response
func (o *GetNodesDefault) WithStatusCode(code int) *GetNodesDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get nodes default response
func (o *GetNodesDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get nodes default response
func (o *GetNodesDefault) WithPayload(payload *models.Error) *GetNodesDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get nodes default response
func (o *GetNodesDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodesDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
