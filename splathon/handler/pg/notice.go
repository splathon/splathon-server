package pg

import (
	"context"
	"os"
	"time"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
)

func (h *Handler) ListNotices(ctx context.Context, params operations.ListNoticesParams) (*models.ListNoticesResponse, error) {
	if _, err := h.getTokenSession(params.XSPLATHONAPITOKEN); err != nil {
		return nil, err
	}
	// TODO(haya14busa): remove.
	if os.Getenv("SPLATHON_DUMMY_NOTICE") == "1" {
		return dummyNotices(), nil
	}
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	var ns []*Notice
	if err := h.db.Where("event_id = ?", eventID).Order("created_at DESC").Find(&ns).Error; err != nil {
		return nil, err
	}

	res := &models.ListNoticesResponse{
		Notices: make([]*models.Notice, len(ns)),
	}
	for i, n := range ns {
		res.Notices[i] = convertNotice(n)
	}
	return res, nil
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
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return err
	}
	n := &Notice{
		EventId:   eventID,
		Text:      *params.Notice.Text,
		CreatedAt: time.Unix(*params.Notice.TimestampSec, 0),
		UpdatedAt: time.Now(),
	}
	if params.Notice.ID != 0 {
		n.Id = params.Notice.ID
		h.db.Save(&n)
	} else {
		h.db.Create(&n)
	}
	return h.db.Where(&Notice{Id: params.Notice.ID}).Assign(&n).FirstOrCreate(&Notice{}).Error
}

func (h *Handler) DeleteNotice(ctx context.Context, params admin.DeleteNoticeParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}
	return h.db.Where("id = ?", params.NoticeID).Delete(&Notice{}).Error
}
