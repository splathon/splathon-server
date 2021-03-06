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

// Ranking 予選ランキング
// swagger:model Ranking
type Ranking struct {

	// ランキング計算時点の説明。e.g. 予選第2ラウンド終了時
	RankTime string `json:"rank_time,omitempty"`

	// rankings
	Rankings []*Rank `json:"rankings"`
}

// Validate validates this ranking
func (m *Ranking) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRankings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ranking) validateRankings(formats strfmt.Registry) error {

	if swag.IsZero(m.Rankings) { // not required
		return nil
	}

	for i := 0; i < len(m.Rankings); i++ {
		if swag.IsZero(m.Rankings[i]) { // not required
			continue
		}

		if m.Rankings[i] != nil {
			if err := m.Rankings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rankings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Ranking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ranking) UnmarshalBinary(b []byte) error {
	var res Ranking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
