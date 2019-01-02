// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"mime/multipart"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadAppValuesParams creates a new UploadAppValuesParams object
// no default values defined in spec.
func NewUploadAppValuesParams() UploadAppValuesParams {

	return UploadAppValuesParams{}
}

// UploadAppValuesParams contains all the bound params for the upload app values operation
// typically these are obtained from a http.Request
//
// swagger:parameters uploadAppValues
type UploadAppValuesParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*App name
	  In: formData
	*/
	AppName *string
	/*Cluster name to deploy app
	  In: formData
	*/
	ClusterName *string
	/*Owner of the app
	  In: formData
	*/
	Owner *string
	/*App values.yaml to upload
	  In: formData
	*/
	Values io.ReadCloser
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUploadAppValuesParams() beforehand.
func (o *UploadAppValuesParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		if err != http.ErrNotMultipart {
			return errors.New(400, "%v", err)
		} else if err := r.ParseForm(); err != nil {
			return errors.New(400, "%v", err)
		}
	}
	fds := runtime.Values(r.Form)

	fdAppName, fdhkAppName, _ := fds.GetOK("appName")
	if err := o.bindAppName(fdAppName, fdhkAppName, route.Formats); err != nil {
		res = append(res, err)
	}

	fdClusterName, fdhkClusterName, _ := fds.GetOK("clusterName")
	if err := o.bindClusterName(fdClusterName, fdhkClusterName, route.Formats); err != nil {
		res = append(res, err)
	}

	fdOwner, fdhkOwner, _ := fds.GetOK("owner")
	if err := o.bindOwner(fdOwner, fdhkOwner, route.Formats); err != nil {
		res = append(res, err)
	}

	values, valuesHeader, err := r.FormFile("values")
	if err != nil && err != http.ErrMissingFile {
		res = append(res, errors.New(400, "reading file %q failed: %v", "values", err))
	} else if err == http.ErrMissingFile {
		// no-op for missing but optional file parameter
	} else if err := o.bindValues(values, valuesHeader); err != nil {
		res = append(res, err)
	} else {
		o.Values = &runtime.File{Data: values, Header: valuesHeader}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAppName binds and validates parameter AppName from formData.
func (o *UploadAppValuesParams) bindAppName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.AppName = &raw

	return nil
}

// bindClusterName binds and validates parameter ClusterName from formData.
func (o *UploadAppValuesParams) bindClusterName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.ClusterName = &raw

	return nil
}

// bindOwner binds and validates parameter Owner from formData.
func (o *UploadAppValuesParams) bindOwner(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Owner = &raw

	return nil
}

// bindValues binds file parameter Values.
//
// The only supported validations on files are MinLength and MaxLength
func (o *UploadAppValuesParams) bindValues(file multipart.File, header *multipart.FileHeader) error {
	return nil
}