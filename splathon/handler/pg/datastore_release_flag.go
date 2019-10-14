package pg

import (
	"context"

	"cloud.google.com/go/datastore"
	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
)

type QualifierRelease struct {
	EventID        int64
	QualifierRound int32
}

func (q *QualifierRelease) dsKey(ctx context.Context) *datastore.Key {
	const kind = "QualifierRelease"
	return datastore.IDKey(kind, q.EventID, nil)
}

func UpdateQualifierRelease(ctx context.Context, eventID int64, qualifierRound int32) error {
	e := &QualifierRelease{
		EventID:        eventID,
		QualifierRound: qualifierRound,
	}
	cli, err := datastoreClient(ctx)
	if err != nil {
		return err
	}
	_, err = cli.Put(ctx, e.dsKey(ctx), e)
	return err
}

func GetQualifierRelease(ctx context.Context, eventID int64) (int32, error) {
	k := (&QualifierRelease{EventID: eventID}).dsKey(ctx)
	e := new(QualifierRelease)
	cli, err := datastoreClient(ctx)
	if err != nil {
		return 0, err
	}
	if err := cli.Get(ctx, k, e); err != nil {
		if err == datastore.ErrNoSuchEntity {
			return -1, nil
		}
		return 0, err
	}
	return e.QualifierRound, nil
}

func (h *Handler) GetReleaseQualifier(ctx context.Context, params admin.GetReleaseQualifierParams) (int32, error) {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return 0, err
	}
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return 0, err
	}
	return GetQualifierRelease(ctx, eventID)
}

func (h *Handler) UpdateReleaseQualifier(ctx context.Context, params admin.UpdateReleaseQualifierParams) error {
	if err := h.checkAdminAuth(params.XSPLATHONAPITOKEN); err != nil {
		return err
	}
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return err
	}
	defer h.clearResultCache(eventID)
	return UpdateQualifierRelease(ctx, eventID, params.Request.Round)
}
