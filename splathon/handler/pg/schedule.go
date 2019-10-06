package pg

import (
	"context"
	"time"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) GetSchedule(context.Context, operations.GetScheduleParams) (*models.Schedule, error) {
	return &models.Schedule{
		// Return dummy data.
		// https://url-c.com/tc/
		Entries: []*models.ScheduleEntry{
			{
				Title:             "開会",
				StartTimestampSec: 1571526000,
			},
			{
				Title:             "DJ #1 開始",
				StartTimestampSec: 1571530500,
				DurationSec:       int64((30 * time.Minute).Seconds()),
			},
			{
				Title:             "予選第1ラウンド　発表",
				StartTimestampSec: 1571532300,
			},
			{
				Title:             "開会式",
				StartTimestampSec: 1571533200,
				DurationSec:       int64((20 * time.Minute).Seconds()),
			},
			{
				Title:             "予選第1ラウンド",
				StartTimestampSec: 1571534400,
				DurationSec:       int64((88 * time.Minute).Seconds()),
			},
			{
				Title:             "予選第2ラウンド",
				StartTimestampSec: 1571540280,
				DurationSec:       int64((68 * time.Minute).Seconds()),
			},
			{
				Title:             "予選第3ラウンド",
				StartTimestampSec: 1571547060,
				DurationSec:       int64((68 * time.Minute).Seconds()),
			},
			{
				Title:             "予選第4ラウンド",
				StartTimestampSec: 1571551440,
				DurationSec:       int64((68 * time.Minute).Seconds()),
			},
			{
				Title:             "決勝T",
				StartTimestampSec: 1571556120,
				DurationSec:       int64((100 * time.Minute).Seconds()),
			},
			{
				Title:             "閉会式",
				StartTimestampSec: 1571563020,
				DurationSec:       int64((10 * time.Minute).Seconds()),
			},
		},
	}, nil
}
