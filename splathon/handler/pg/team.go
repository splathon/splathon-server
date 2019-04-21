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
		// TODO(haya14busa): fill in real members.
		fillInDummyMembers(false, r.Teams[i])
	}
	return r, nil
}

func (h *Handler) GetTeamDetail(ctx context.Context, params operations.GetTeamDetailParams) (*models.Team, error) {
	var t Team
	if err := h.db.Where("id = ?", params.TeamID).Find(&t).Error; err != nil {
		return nil, err
	}
	team := convertTeam(&t)
	// TODO(haya14busa): fill in real members with detail data.
	fillInDummyMembers(true, team)
	return team, nil
}
