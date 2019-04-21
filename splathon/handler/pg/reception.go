package pg

import (
	"context"
	"fmt"
	"net/http"
	"os"

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
	if token.IsAdmin && token.SlackUserID == "" {
		slackUserID = "admin" // HACK for testing with admin account.
	}

	if slackUserID == "" {
		return nil, &serror.Error{
			Code:    http.StatusUnauthorized,
			Message: "login user doesn't have associated slack user ID.",
		}
	}
	resp := &models.ReceptionResponse{
		Building: &models.ReceptionCode{
			Name: "ビル入館コード",
			Description: `QRコードを入退ゲートにかざしご入館・退館ください。
同伴者様がいる場合は同伴者様にもアプリをインストールしていただくか、Splathon 運営に連絡の上事前にQRコードを印刷してご持参ください。
`,
			ShortText: "TODO(haya14busa): こっちに名前、参加費、同伴者(ないし運営チェック)の有無など簡潔に書いてもいいかも。",
			CodeType:  models.ReceptionCodeCodeTypeQrcode,
			Code:      slackUserID,
			QrcodeImg: googleQRCodeImageURL(slackUserID),
		},
		Splathon: &models.ReceptionCode{
			Name: "会場入場コード",
			Description: `Splathon 会場でこのQRコードを表示して受付してください。
TODO(haya14busa): ここに最低限の参加者情報や払うべき金額を事前に表示する。
`,
			ShortText: fmt.Sprintf("来客用入館証 (受付番号：%s)", os.Getenv("SPLATHON_BUILDING_RECEPTION_NUMBER")),
			CodeType:  models.ReceptionCodeCodeTypeQrcode,
			Code:      os.Getenv("SPLATHON_BUILDING_CODE"),
			QrcodeImg: os.Getenv("SPLATHON_BUILDING_QRCODE_URL"),
		},
	}
	return resp, nil
}

func googleQRCodeImageURL(code string) string {
	return fmt.Sprintf("https://chart.apis.google.com/chart?chs=142x142&cht=qr&chl=%s", code)
}
