package pg

import (
	"context"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/splathon/spldata"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) GetEvent(ctx context.Context, params operations.GetEventParams) (*models.Event, error) {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}
	var e Event
	if err := h.db.Where("id = ?", eventID).Find(&e).Error; err != nil {
		return nil, err
	}
	event := &models.Event{
		Numbering: swag.Int32(int32(eventID)),
		Name:      swag.String(e.Name),
		Rules:     make([]*models.Rule, 0),
		Stages:    make([]*models.Stage, 0),
	}
	for _, r := range spldata.ListRules() {
		event.Rules = append(event.Rules, convertRule(r))
	}
	for _, s := range spldata.ListStages() {
		event.Stages = append(event.Stages, convertStage(s))
	}
	return event, nil
}

func convertRule(r spldata.Rule) *models.Rule {
	return &models.Rule{Key: swag.String(r.Key), Name: r.Name}
}

func convertStage(s spldata.Stage) *models.Stage {
	return &models.Stage{ID: swag.Int32(int32(s.ID)), Name: s.Name}
}
