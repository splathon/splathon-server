package pg

import (
	"context"
	"fmt"

	"github.com/splathon/splathon-server/swagger/restapi/operations/admin"
	"google.golang.org/appengine/datastore"
)

var _ = datastore.NewKey

type QualifierRelease struct {
	EventID        int64
	QualifierRound int32
}

func (q *QualifierRelease) dsKey(ctx context.Context) *datastore.Key {
	const kind = "QualifierRelease"
	return datastore.NewKey(ctx, kind, "", q.EventID, nil)
}

func UpdateQualifierRelease(ctx context.Context, eventID int64, qualifierRound int32) error {
	e := &QualifierRelease{
		EventID:        eventID,
		QualifierRound: qualifierRound,
	}
	_, err := datastore.Put(ctx, e.dsKey(ctx), e)
	return err
}

func GetQualifierRelease(ctx context.Context, eventID int64) (int32, error) {
	k := (&QualifierRelease{EventID: eventID}).dsKey(ctx)
	e := new(QualifierRelease)
	if err := datastore.Get(ctx, k, e); err != nil {
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
	fmt.Printf("remove all result cache")
	h.resultCacheMu.Lock()
	defer h.resultCacheMu.Unlock()
	h.resultCache = make(map[resultCacheKey]*resultCache)
	return UpdateQualifierRelease(ctx, eventID, params.Request.Round)
}
