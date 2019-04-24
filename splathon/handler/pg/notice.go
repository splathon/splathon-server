package pg

import (
	"context"
	"time"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) ListNotices(ctx context.Context, params operations.ListNoticesParams) (*models.ListNoticesResponse, error) {
	if _, err := h.getTokenSession(params.XSPLATHONAPITOKEN); err != nil {
		return nil, err
	}
	return dummyNotices(), nil
}

func dummyNotices() *models.ListNoticesResponse {
	now := time.Now()
	return &models.ListNoticesResponse{
		Notices: []*models.Notice{
			{
				Text: swag.String(`(ダミーテキスト) いよいよ明日はSplathon#10ですね！むっちゃドキドキしてきた。。。
みなさんも今日くらいは練習休んで明日に備えますよね？？？`),
				TimestampSec: swag.Int64(now.Unix()),
			},
			{
				Text: swag.String(`## Splathon前夜祭　特別メニューのお知らせ

Splathon前夜祭では、参加費は別にお金を払うことによって食べられる特別メニュー「廉価版5桁丼」をご用意します。
メニュー詳細は以下の記事をご参照ください`),
				TimestampSec: swag.Int64(now.Add(-time.Minute).Unix()),
			},
			{
				Text: swag.String(`コミュニティチームより、Splathon #10当日のガイドラインが公開されました。
こちらもしおりと併せてご一読ください。`),
				TimestampSec: swag.Int64(now.Add(-2 * time.Hour).Unix()),
			},
			{
				Text: swag.String(`splathon#10サイトと、併せて『しおり』公開してます！！
しおりは参加者全員、一読お願いしますっ 🙏`),
				TimestampSec: swag.Int64(now.Add(-25 * time.Hour).Unix()),
			},
		},
	}
}
