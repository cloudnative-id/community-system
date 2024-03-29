// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Speaker speaker
//
// swagger:model Speaker
type Speaker struct {

	// company
	// Required: true
	Company *string `json:"company"`

	// image
	// Required: true
	Image *string `json:"image"`

	// name
	// Required: true
	Name *string `json:"name"`

	// position
	// Required: true
	Position *string `json:"position"`

	// title
	// Required: true
	Title *string `json:"title"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this speaker
func (m *Speaker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompany(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Speaker) validateCompany(formats strfmt.Registry) error {

	if err := validate.Required("company", "body", m.Company); err != nil {
		return err
	}

	return nil
}

func (m *Speaker) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *Speaker) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Speaker) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	return nil
}

func (m *Speaker) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this speaker based on context it is used
func (m *Speaker) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Speaker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Speaker) UnmarshalBinary(b []byte) error {
	var res Speaker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
