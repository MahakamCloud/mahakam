// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Metadata metadata
// swagger:model metadata
type Metadata struct {

	// ssh keys
	SSHKeys []string `json:"sshKeys"`

	// userdata
	Userdata string `json:"userdata,omitempty"`

	// metadata
	Metadata map[string]string `json:"-"`
}

// UnmarshalJSON unmarshals this object with additional properties from JSON
func (m *Metadata) UnmarshalJSON(data []byte) error {
	// stage 1, bind the properties
	var stage1 struct {

		// ssh keys
		SSHKeys []string `json:"sshKeys"`

		// userdata
		Userdata string `json:"userdata,omitempty"`
	}
	if err := json.Unmarshal(data, &stage1); err != nil {
		return err
	}
	var rcv Metadata

	rcv.SSHKeys = stage1.SSHKeys

	rcv.Userdata = stage1.Userdata

	*m = rcv

	// stage 2, remove properties and add to map
	stage2 := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &stage2); err != nil {
		return err
	}

	delete(stage2, "sshKeys")

	delete(stage2, "userdata")

	// stage 3, add additional properties values
	if len(stage2) > 0 {
		result := make(map[string]string)
		for k, v := range stage2 {
			var toadd string
			if err := json.Unmarshal(v, &toadd); err != nil {
				return err
			}
			result[k] = toadd
		}
		m.Metadata = result
	}

	return nil
}

// MarshalJSON marshals this object with additional properties into a JSON object
func (m Metadata) MarshalJSON() ([]byte, error) {
	var stage1 struct {

		// ssh keys
		SSHKeys []string `json:"sshKeys"`

		// userdata
		Userdata string `json:"userdata,omitempty"`
	}

	stage1.SSHKeys = m.SSHKeys

	stage1.Userdata = m.Userdata

	// make JSON object for known properties
	props, err := json.Marshal(stage1)
	if err != nil {
		return nil, err
	}

	if len(m.Metadata) == 0 {
		return props, nil
	}

	// make JSON object for the additional properties
	additional, err := json.Marshal(m.Metadata)
	if err != nil {
		return nil, err
	}

	if len(props) < 3 {
		return additional, nil
	}

	// concatenate the 2 objects
	props[len(props)-1] = ','
	return append(props, additional[1:]...), nil
}

// Validate validates this metadata
func (m *Metadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metadata) UnmarshalBinary(b []byte) error {
	var res Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
