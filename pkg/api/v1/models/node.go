// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Node node
// swagger:model node
type Node struct {
	BaseResource

	// metadata
	Metadata *NodeMetadata `json:"metadata,omitempty"`

	// network configs
	NetworkConfigs []*NodeNetworkConfig `json:"networkConfigs"`

	// spec
	Spec *NodeSpec `json:"spec,omitempty"`

	// status
	Status *NodeStatus `json:"status,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Node) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseResource = aO0

	// now for regular properties
	var propsNode struct {
		Metadata *NodeMetadata `json:"metadata,omitempty"`

		NetworkConfigs []*NodeNetworkConfig `json:"networkConfigs"`

		Spec *NodeSpec `json:"spec,omitempty"`

		Status *NodeStatus `json:"status,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsNode); err != nil {
		return err
	}
	m.Metadata = propsNode.Metadata

	m.NetworkConfigs = propsNode.NetworkConfigs

	m.Spec = propsNode.Spec

	m.Status = propsNode.Status

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Node) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BaseResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsNode struct {
		Metadata *NodeMetadata `json:"metadata,omitempty"`

		NetworkConfigs []*NodeNetworkConfig `json:"networkConfigs"`

		Spec *NodeSpec `json:"spec,omitempty"`

		Status *NodeStatus `json:"status,omitempty"`
	}
	propsNode.Metadata = m.Metadata

	propsNode.NetworkConfigs = m.NetworkConfigs

	propsNode.Spec = m.Spec

	propsNode.Status = m.Status

	jsonDataPropsNode, errNode := swag.WriteJSON(propsNode)
	if errNode != nil {
		return nil, errNode
	}
	_parts = append(_parts, jsonDataPropsNode)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseResource
	if err := m.BaseResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkConfigs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpec(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *Node) validateNetworkConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkConfigs) { // not required
		return nil
	}

	for i := 0; i < len(m.NetworkConfigs); i++ {
		if swag.IsZero(m.NetworkConfigs[i]) { // not required
			continue
		}

		if m.NetworkConfigs[i] != nil {
			if err := m.NetworkConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networkConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Node) validateSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.Spec) { // not required
		return nil
	}

	if m.Spec != nil {
		if err := m.Spec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spec")
			}
			return err
		}
	}

	return nil
}

func (m *Node) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
