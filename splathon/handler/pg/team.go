package pg

import (
	"context"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) ListTeams(ctx context.Context, params operations.ListTeamsParams) (*models.Teams, error) {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	var teams []*Team
	if err := h.db.Where("event_id = ?", eventID).Order("id asc").Find(&teams).Error; err != nil {
		return nil, err
	}

	r := &models.Teams{
		Teams: make([]*models.Team, len(teams)),
	}
	for i, t := range teams {
		r.Teams[i] = convertTeam(t)
	}
	return r, nil
}
