package pg

import (
	"context"

	"github.com/splathon/splathon-server/swagger/models"
	"google.golang.org/appengine/datastore"
)

type EventSchedule struct {
	EventID  int64
	Schedule *models.Schedule
}

func (es *EventSchedule) dsKey(ctx context.Context) *datastore.Key {
	const kind = "EventSchedule"
	return datastore.NewKey(ctx, kind, "", es.EventID, nil)
}

func UpdateEventSchedule(ctx context.Context, eventID int64, schedule *models.Schedule) error {
	e := &EventSchedule{
		EventID:  eventID,
		Schedule: schedule,
	}
	_, err := datastore.Put(ctx, e.dsKey(ctx), e)
	return err
}

func GetEventSchedule(ctx context.Context, eventID int64) (*models.Schedule, error) {
	k := (&EventSchedule{EventID: eventID}).dsKey(ctx)
	e := new(EventSchedule)
	if err := datastore.Get(ctx, k, e); err != nil {
		return nil, err
	}
	return e.Schedule, nil
}
