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

// UpdateReceptionRequest update reception request
// swagger:model UpdateReceptionRequest
type UpdateReceptionRequest struct {

	// complete
	// Required: true
	Complete *bool `json:"complete"`

	// participant
	// Required: true
	Participant *ParticipantReception `json:"participant"`
}

// Validate validates this update reception request
func (m *UpdateReceptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipant(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateReceptionRequest) validateComplete(formats strfmt.Registry) error {

	if err := validate.Required("complete", "body", m.Complete); err != nil {
		return err
	}

	return nil
}

func (m *UpdateReceptionRequest) validateParticipant(formats strfmt.Registry) error {

	if err := validate.Required("participant", "body", m.Participant); err != nil {
		return err
	}

	if m.Participant != nil {
		if err := m.Participant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("participant")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateReceptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateReceptionRequest) UnmarshalBinary(b []byte) error {
	var res UpdateReceptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
