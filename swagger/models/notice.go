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

// Notice notice
// swagger:model Notice
type Notice struct {

	// text
	// Required: true
	Text *string `json:"text"`

	// timestamp sec
	// Required: true
	TimestampSec *int64 `json:"timestamp_sec"`
}

// Validate validates this notice
func (m *Notice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestampSec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notice) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

func (m *Notice) validateTimestampSec(formats strfmt.Registry) error {

	if err := validate.Required("timestamp_sec", "body", m.TimestampSec); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Notice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Notice) UnmarshalBinary(b []byte) error {
	var res Notice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
