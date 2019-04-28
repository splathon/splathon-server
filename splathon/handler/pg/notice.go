package pg

import (
	"context"
	"errors"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
)

func (h *Handler) ListNotices(ctx context.Context, params operations.ListNoticesParams) (*models.ListNoticesResponse, error) {
	if _, err := h.getTokenSession(params.XSPLATHONAPITOKEN); err != nil {
		return nil, err
	}
	return dummyNotices(), nil
}

func dummyNotices() *models.ListNoticesResponse {
	return &models.ListNoticesResponse{
		Notices: []*models.Notice{
			{
				Text:         swag.String(`すでにしおりは読みましたか？まだの方は最新版をチェック! http://bit.ly/2La6fSR`),
				TimestampSec: swag.Int64(1556460633),
			},
			{
				Text:         swag.String(`Splathon#10がいよいよ5月1日に開催されます!`),
				TimestampSec: swag.Int64(1556453100),
			},
		},
	}
}

func (h *Handler) WriteNotice(ctx context.Context, params admin.WriteNoticeParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}
	return errors.New("operation .WriteNotices has not yet been implemented")
}
