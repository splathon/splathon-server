// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DeleteQualifierRequest delete qualifier request
// swagger:model DeleteQualifierRequest
type DeleteQualifierRequest struct {

	// round
	Round int32 `json:"round,omitempty"`
}

// Validate validates this delete qualifier request
func (m *DeleteQualifierRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteQualifierRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteQualifierRequest) UnmarshalBinary(b []byte) error {
	var res DeleteQualifierRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
