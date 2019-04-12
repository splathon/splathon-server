// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Room Roomごとのマッチ
// swagger:model Room
type Room struct {

	// Room ID.
	ID int32 `json:"id,omitempty"`

	// matches
	// Required: true
	Matches []*Match `json:"matches"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this room
func (m *Room) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMatches(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Room) validateMatches(formats strfmt.Registry) error {

	if err := validate.Required("matches", "body", m.Matches); err != nil {
		return err
	}

	for i := 0; i < len(m.Matches); i++ {
		if swag.IsZero(m.Matches[i]) { // not required
			continue
		}

		if m.Matches[i] != nil {
			if err := m.Matches[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("matches" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Room) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Room) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Room) UnmarshalBinary(b []byte) error {
	var res Room
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
