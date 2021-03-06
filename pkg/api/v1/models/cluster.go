// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cluster cluster
// swagger:model cluster
type Cluster struct {

	// component failures
	ComponentFailures []string `json:"componentFailures"`

	// gre key
	GreKey int64 `json:"greKey,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// name
	// Required: true
	// Min Length: 1
	Name *string `json:"name"`

	// node failures
	NodeFailures []string `json:"nodeFailures"`

	// node size
	NodeSize *string `json:"nodeSize,omitempty"`

	// num nodes
	// Maximum: 10
	// Minimum: 1
	NumNodes int64 `json:"numNodes,omitempty"`

	// owner
	// Min Length: 1
	Owner string `json:"owner,omitempty"`

	// pod failures
	PodFailures []string `json:"podFailures"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumNodes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateNumNodes(formats strfmt.Registry) error {

	if swag.IsZero(m.NumNodes) { // not required
		return nil
	}

	if err := validate.MinimumInt("numNodes", "body", int64(m.NumNodes), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("numNodes", "body", int64(m.NumNodes), 10, false); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if err := validate.MinLength("owner", "body", string(m.Owner), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
