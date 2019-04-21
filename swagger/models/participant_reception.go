// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ParticipantReception participant reception
// swagger:model ParticipantReception
type ParticipantReception struct {

	// 所属企業名
	CompanyName string `json:"company_name,omitempty"`

	// カタカナのフルネーム。 e.g. ヤマダタロウ
	FullnameKana string `json:"fullname_kana,omitempty"`

	// 同伴者がいるかどうか。いる場合は用スプレッドシート確認。
	HasCompanion bool `json:"has_companion,omitempty"`

	// Nintendo Switch doc を持ってきたか
	HasSwitchDock bool `json:"has_switch_dock,omitempty"`

	// playerとして参加するかどうか。falseならスタッフか観戦
	IsPlayer bool `json:"is_player,omitempty"`

	// スタッフかどうか
	IsStaff bool `json:"is_staff,omitempty"`

	// 懇親会に参加するか否か
	JoinParty bool `json:"join_party,omitempty"`

	// ハンドルネーム。 e.g. みーくん
	Nickname string `json:"nickname,omitempty"`

	// 合計参加費(円)
	ParticipantFee int32 `json:"participant_fee,omitempty"`

	// チームID(一応)
	TeamID int64 `json:"team_id,omitempty"`

	// チーム名
	TeamName string `json:"team_name,omitempty"`
}

// Validate validates this participant reception
func (m *ParticipantReception) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ParticipantReception) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ParticipantReception) UnmarshalBinary(b []byte) error {
	var res ParticipantReception
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}