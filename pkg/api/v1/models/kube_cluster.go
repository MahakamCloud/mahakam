// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// KubeCluster kube cluster
// swagger:model kubeCluster
type KubeCluster struct {
	BaseResource

	// node size
	// Enum: [xs s m l xl]
	NodeSize *string `json:"nodeSize,omitempty"`

	// num nodes
	// Maximum: 10
	// Minimum: 1
	NumNodes int64 `json:"numNodes,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *KubeCluster) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BaseResource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BaseResource = aO0

	// now for regular properties
	var propsKubeCluster struct {
		NodeSize *string `json:"nodeSize,omitempty"`

		NumNodes int64 `json:"numNodes,omitempty"`
	}
	if err := swag.ReadJSON(raw, &propsKubeCluster); err != nil {
		return err
	}
	m.NodeSize = propsKubeCluster.NodeSize

	m.NumNodes = propsKubeCluster.NumNodes

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m KubeCluster) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.BaseResource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	// now for regular properties
	var propsKubeCluster struct {
		NodeSize *string `json:"nodeSize,omitempty"`

		NumNodes int64 `json:"numNodes,omitempty"`
	}
	propsKubeCluster.NodeSize = m.NodeSize

	propsKubeCluster.NumNodes = m.NumNodes

	jsonDataPropsKubeCluster, errKubeCluster := swag.WriteJSON(propsKubeCluster)
	if errKubeCluster != nil {
		return nil, errKubeCluster
	}
	_parts = append(_parts, jsonDataPropsKubeCluster)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this kube cluster
func (m *KubeCluster) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BaseResource
	if err := m.BaseResource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeSize(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumNodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var kubeClusterTypeNodeSizePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["xs","s","m","l","xl"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		kubeClusterTypeNodeSizePropEnum = append(kubeClusterTypeNodeSizePropEnum, v)
	}
}

const (

	// KubeClusterNodeSizeXs captures enum value "xs"
	KubeClusterNodeSizeXs string = "xs"

	// KubeClusterNodeSizeS captures enum value "s"
	KubeClusterNodeSizeS string = "s"

	// KubeClusterNodeSizeM captures enum value "m"
	KubeClusterNodeSizeM string = "m"

	// KubeClusterNodeSizeL captures enum value "l"
	KubeClusterNodeSizeL string = "l"

	// KubeClusterNodeSizeXl captures enum value "xl"
	KubeClusterNodeSizeXl string = "xl"
)

// prop value enum
func (m *KubeCluster) validateNodeSizeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, kubeClusterTypeNodeSizePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *KubeCluster) validateNodeSize(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeSize) { // not required
		return nil
	}

	// value enum
	if err := m.validateNodeSizeEnum("nodeSize", "body", *m.NodeSize); err != nil {
		return err
	}

	return nil
}

func (m *KubeCluster) validateNumNodes(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *KubeCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KubeCluster) UnmarshalBinary(b []byte) error {
	var res KubeCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
