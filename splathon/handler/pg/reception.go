package pg

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/splathon/splathon-server/splathon/serror"
	"github.com/splathon/splathon-server/swagger/models"
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

	var ps []*Participant
	if err := h.db.Where("slack_user_id = ?", slackUserID).Find(&ps).Error; err != nil || len(ps) == 0 {
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

func googleQRCodeImageURL(code string) string {
	return fmt.Sprintf("https://chart.apis.google.com/chart?chs=142x142&cht=qr&chl=%s", code)
}

func boolToJapanese(b bool) string {
	if b {
		return "あり"
	}
	return "なし"
}
