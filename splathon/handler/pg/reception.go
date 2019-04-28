package pg

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-openapi/swag"
	"github.com/jinzhu/gorm"
	"github.com/splathon/splathon-server/splathon/serror"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
	"github.com/splathon/splathon-server/swagger/restapi/operations/reception"
)

func (h *Handler) GetReception(ctx context.Context, params reception.GetReceptionParams) (*models.ReceptionResponse, error) {
	token, err := h.getTokenSession(params.XSPLATHONAPITOKEN)
	if err != nil {
		return nil, err
	}
	slackUserID := token.SlackUserID
	if slackUserID == "" {
		return nil, &serror.Error{
			Code:    http.StatusUnauthorized,
			Message: "login user doesn't have associated slack user ID.",
		}
	}

	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	var ps []*Participant
	if err := h.db.Where("event_id = ?  AND slack_user_id = ?", eventID, slackUserID).Find(&ps).Error; err != nil || len(ps) == 0 {
		return nil, fmt.Errorf("invalid token: participant not found with id=%q", slackUserID)
	}

	sumFee := 0
	nicknames := make([]string, len(ps))
	hasCompanion := false
	joinParty := false
	for i, p := range ps {
		sumFee += int(p.Fee)
		nicknames[i] = p.Nickname
		hasCompanion = hasCompanion || p.HasCompanion
		joinParty = joinParty || p.JoinParty
	}
	thonShortData := fmt.Sprintf("[%s] 合計支払い金額: %d円 (懇親会参加: %s)", strings.Join(nicknames, ","), sumFee, boolToJapanese(joinParty))

	if hasCompanion {
		thonShortData = fmt.Sprintf("[%s] 参考支払い金額: %d円, 同伴者様の懇親会参加の有無などで金額が前後するので受付でお申し付けください。", strings.Join(nicknames, ","), sumFee)
	}

	resp := &models.ReceptionResponse{
		Building: &models.ReceptionCode{
			Name: "ビル入館コード",
			Description: `QRコードを入退ゲートにかざしご入館・退館ください。
同伴者様がいる場合は同伴者様にもアプリをインストールしていただくか、Splathon 運営に連絡の上事前にQRコードを印刷してご持参ください。
`,
			ShortText: fmt.Sprintf("来客用入館証 (受付番号：%s)", os.Getenv("SPLATHON_BUILDING_RECEPTION_NUMBER")),
			CodeType:  models.ReceptionCodeCodeTypeQrcode,
			Code:      os.Getenv("SPLATHON_BUILDING_CODE"),
			QrcodeImg: os.Getenv("SPLATHON_BUILDING_QRCODE_URL"),
		},
		Splathon: &models.ReceptionCode{
			Name:        "会場入場コード",
			Description: `Splathon 会場でこのQRコードを表示して受付してください。`,
			ShortText:   thonShortData,
			CodeType:    models.ReceptionCodeCodeTypeQrcode,

			Code:      slackUserID,
			QrcodeImg: googleQRCodeImageURL(slackUserID),
		},
	}
	return resp, nil
}

func (h *Handler) GetParticipantsDataForReception(ctx context.Context, params operations.GetParticipantsDataForReceptionParams) (*models.ReceptionPartcipantsDataResponse, error) {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return nil, err
	}

	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	slackID := params.SplathonReceptionCode

	ps, err := h.participantsByReceptionCode(eventID, params.SplathonReceptionCode)
	if err != nil {
		return nil, err
	}

	response := &models.ReceptionPartcipantsDataResponse{
		Description:     "同伴者がいる場合は別途スプレッドシートを参照してください。",
		SLACKInternalID: slackID,
		Participants:    make([]*models.ParticipantReception, len(ps)),
	}
	for i, p := range ps {
		r := convertParticipant(p)
		if p.TeamId.Valid {
			var team Team
			if err := h.db.Select("name").Where("id = ?", p.TeamId.Int64).Find(&team).Error; err != nil {
				return nil, err
			}
			r.TeamID = p.TeamId.Int64
			r.TeamName = team.Name
		}
		response.Participants[i] = r
	}
	return response, nil
}

func (h *Handler) CompleteReception(ctx context.Context, params operations.CompleteReceptionParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}

	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return err
	}

	ps, err := h.participantsByReceptionCode(eventID, params.SplathonReceptionCode)
	if err != nil {
		return err
	}

	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	for _, p := range ps {
		r := &Reception{
			ParticipantId: p.Id,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		}
		var res Reception
		if err := tx.Where(Reception{ParticipantId: p.Id}).Assign(&r).FirstOrCreate(&res).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit().Error
}

func (h *Handler) ListReception(ctx context.Context, params admin.ListReceptionParams) (*models.ListReceptionResponse, error) {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return nil, err
	}
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}
	var ps []*Participant
	if err := h.db.Where("event_id = ?", eventID).Find(&ps).Error; err != nil {
		return nil, err
	}
	ids := make([]int64, len(ps))
	for i, p := range ps {
		ids[i] = p.Id
	}
	var rs []*Reception
	if err := h.db.Where("participant_id IN (?)", ids).Find(&rs).Error; err != nil {
		return nil, err
	}
	id2r := make(map[int64]*Reception)
	for _, r := range rs {
		id2r[r.ParticipantId] = r
	}
	resp := &models.ListReceptionResponse{
		Participants: make([]*models.ParticipantReception, len(ps)),
	}
	for i, p := range ps {
		participant := convertParticipant(p)
		if r, ok := id2r[p.Id]; ok {
			participant.Reception = &models.Reception{
				ID:                    swag.Int64(r.Id),
				ParticipantID:         swag.Int64(p.Id),
				Memo:                  r.Memo,
				CreatedAtTimestampSec: swag.Int64(r.CreatedAt.Unix()),
				UpdatedAtTimestampSec: swag.Int64(r.UpdatedAt.Unix()),
			}
		}
		resp.Participants[i] = participant
	}
	return resp, nil
}

func (h *Handler) UpdateReception(ctx context.Context, params admin.UpdateReceptionParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}
	tx := h.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := updateReception(ctx, tx, params); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func updateReception(ctx context.Context, tx *gorm.DB, params admin.UpdateReceptionParams) error {
	participantID := *params.Request.Participant.ID

	// Update reception record.
	if *params.Request.Complete {
		r := &Reception{
			ParticipantId: participantID,
			Memo:          params.Request.Participant.Reception.Memo,
			UpdatedAt:     time.Now(),
		}
		if params.Request.Participant.Reception == nil {
			r.CreatedAt = time.Now()
		}
		var res Reception
		if err := tx.Where(Reception{ParticipantId: participantID}).Assign(&r).FirstOrCreate(&res).Error; err != nil {
			return err
		}
	} else {
		tx.Where("participant_id = ?", participantID).Delete(&Reception{})
	}

	// Update participant data.
	p := params.Request.Participant
	return tx.Model(&Participant{Id: participantID}).Updates(map[string]interface{}{
		"fee":             *p.ParticipantFee,
		"has_companion":   *p.HasCompanion,
		"join_party":      *p.JoinParty,
		"is_staff":        *p.IsStaff,
		"is_player":       *p.IsPlayer,
		"has_switch_dock": *p.HasSwitchDock,
	}).Error
}

func (h *Handler) participantsByReceptionCode(eventID int64, code string) ([]*Participant, error) {
	slackID := code
	var ps []*Participant
	if err := h.db.Where("event_id = ? AND slack_user_id = ?", eventID, slackID).Find(&ps).Error; err != nil {
		return nil, err
	}
	if len(ps) == 0 {
		return nil, fmt.Errorf("participants not found (code=%q)", slackID)
	}
	return ps, nil
}

func googleQRCodeImageURL(code string) string {
	return fmt.Sprintf("https://chart.apis.google.com/chart?chs=142x142&cht=qr&chl=%s", code)
}

func boolToJapanese(b bool) string {
	if b {
		return "あり"
	}
	return "なし"
}
