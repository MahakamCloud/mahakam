// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/mahakamcloud/mahakam/pkg/api/v1/models"
)

// UploadAppValuesReader is a Reader for the UploadAppValues structure.
type UploadAppValuesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadAppValuesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUploadAppValuesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUploadAppValuesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadAppValuesCreated creates a UploadAppValuesCreated with default headers values
func NewUploadAppValuesCreated() *UploadAppValuesCreated {
	return &UploadAppValuesCreated{}
}

/*UploadAppValuesCreated handles this case with default header values.

App values added
*/
type UploadAppValuesCreated struct {
}

func (o *UploadAppValuesCreated) Error() string {
	return fmt.Sprintf("[POST /apps/values][%d] uploadAppValuesCreated ", 201)
}

func (o *UploadAppValuesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadAppValuesDefault creates a UploadAppValuesDefault with default headers values
func NewUploadAppValuesDefault(code int) *UploadAppValuesDefault {
	return &UploadAppValuesDefault{
		_statusCode: code,
	}
}

/*UploadAppValuesDefault handles this case with default header values.

error
*/
type UploadAppValuesDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the upload app values default response
func (o *UploadAppValuesDefault) Code() int {
	return o._statusCode
}

func (o *UploadAppValuesDefault) Error() string {
	return fmt.Sprintf("[POST /apps/values][%d] uploadAppValues default  %+v", o._statusCode, o.Payload)
}

func (o *UploadAppValuesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
