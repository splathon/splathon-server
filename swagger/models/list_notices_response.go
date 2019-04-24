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

// ListNoticesResponse list notices response
// swagger:model ListNoticesResponse
type ListNoticesResponse struct {

	// List of notices ordered by reverse chronological order.
	Notices []*Notice `json:"notices"`
}

// Validate validates this list notices response
func (m *ListNoticesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListNoticesResponse) validateNotices(formats strfmt.Registry) error {

	if swag.IsZero(m.Notices) { // not required
		return nil
	}

	for i := 0; i < len(m.Notices); i++ {
		if swag.IsZero(m.Notices[i]) { // not required
			continue
		}

		if m.Notices[i] != nil {
			if err := m.Notices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListNoticesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListNoticesResponse) UnmarshalBinary(b []byte) error {
	var res ListNoticesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
