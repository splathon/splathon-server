package pg

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/splathon/splathon-server/swagger/models"
)

type EventSchedule struct {
	EventID  int64
	Schedule *models.Schedule
}

func (es *EventSchedule) dsKey() *datastore.Key {
	const kind = "EventSchedule"
	return datastore.IDKey(kind, es.EventID, nil)
}

func UpdateEventSchedule(ctx context.Context, eventID int64, schedule *models.Schedule) error {
	cli, err := datastoreClient(ctx)
	if err != nil {
		return err
	}
	e := &EventSchedule{
		EventID:  eventID,
		Schedule: schedule,
	}
	_, err = cli.Put(ctx, e.dsKey(), e)
	return err
}

func GetEventSchedule(ctx context.Context, eventID int64) (*models.Schedule, error) {
	cli, err := datastoreClient(ctx)
	if err != nil {
		return nil, err
	}
	k := (&EventSchedule{EventID: eventID}).dsKey()
	e := new(EventSchedule)
	if err := cli.Get(ctx, k, e); err != nil {
		return nil, err
	}
	return e.Schedule, nil
}
