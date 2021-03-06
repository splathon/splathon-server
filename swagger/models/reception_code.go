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

// ReceptionCode ビル入館情報/Splathon会場受付情報
// swagger:model ReceptionCode
type ReceptionCode struct {

	// code
	Code string `json:"code,omitempty"`

	// code type
	// Enum: [qrcode barcode]
	CodeType string `json:"code_type,omitempty"`

	// 入場の説明
	Description string `json:"description,omitempty"`

	// Splathon会場入場コード/XXXビル入館コード
	Name string `json:"name,omitempty"`

	// Image URL of QR code
	QrcodeImg string `json:"qrcode_img,omitempty"`

	// コードの説明
	ShortText string `json:"short_text,omitempty"`
}

// Validate validates this reception code
func (m *ReceptionCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var receptionCodeTypeCodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["qrcode","barcode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		receptionCodeTypeCodeTypePropEnum = append(receptionCodeTypeCodeTypePropEnum, v)
	}
}

const (

	// ReceptionCodeCodeTypeQrcode captures enum value "qrcode"
	ReceptionCodeCodeTypeQrcode string = "qrcode"

	// ReceptionCodeCodeTypeBarcode captures enum value "barcode"
	ReceptionCodeCodeTypeBarcode string = "barcode"
)

// prop value enum
func (m *ReceptionCode) validateCodeTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, receptionCodeTypeCodeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReceptionCode) validateCodeType(formats strfmt.Registry) error {

	if swag.IsZero(m.CodeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateCodeTypeEnum("code_type", "body", m.CodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReceptionCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReceptionCode) UnmarshalBinary(b []byte) error {
	var res ReceptionCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
