// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateObject create object
//
// swagger:model CreateObject
type CreateObject struct {

	// status
	Status string `json:"status,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this create object
func (m *CreateObject) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create object based on context it is used
func (m *CreateObject) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateObject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateObject) UnmarshalBinary(b []byte) error {
	var res CreateObject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
