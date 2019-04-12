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

// Round 予選/決勝Tラウンド両方扱う。
// swagger:model Round
type Round struct {

	// ラウンド名。e.g. 予選第1ラウンド, 決勝T1回戦, 決勝戦
	// Required: true
	Name *string `json:"name"`

	// rooms
	Rooms []*Room `json:"rooms"`

	// 何ラウンドか。i.e. 予選第Nラウンド, 決勝T N回戦
	Round int32 `json:"round,omitempty"`
}

// Validate validates this round
func (m *Round) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRooms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Round) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Round) validateRooms(formats strfmt.Registry) error {

	if swag.IsZero(m.Rooms) { // not required
		return nil
	}

	for i := 0; i < len(m.Rooms); i++ {
		if swag.IsZero(m.Rooms[i]) { // not required
			continue
		}

		if m.Rooms[i] != nil {
			if err := m.Rooms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rooms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Round) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Round) UnmarshalBinary(b []byte) error {
	var res Round
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
