package pg

import (
	"context"

	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
	"golang.org/x/sync/errgroup"
)

func (h *Handler) ListTeams(ctx context.Context, params operations.ListTeamsParams) (*models.Teams, error) {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	var (
		teams        []*Team
		participants []*Participant
	)

	var eg errgroup.Group

	eg.Go(func() error {
		return h.db.Where("event_id = ?", eventID).Order("id asc").Find(&teams).Error
	})

	eg.Go(func() error {
		return h.db.Where("event_id = ? AND team_id IS NOT NULL", eventID).Find(&participants).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	team2members := make(map[int64][]*models.Member)
	for _, p := range participants {
		if !p.TeamId.Valid {
			continue
		}
		teamID := p.TeamId.Int64
		if _, ok := team2members[teamID]; !ok {
			team2members[teamID] = make([]*models.Member, 0)
		}
		team2members[teamID] = append(team2members[teamID], convertParticipant2TeamMember(p))
	}

	r := &models.Teams{
		Teams: make([]*models.Team, len(teams)),
	}
	for i, t := range teams {
		r.Teams[i] = convertTeam(t)
		if ms, ok := team2members[t.Id]; ok {
			r.Teams[i].Members = ms
		} else {
			// TODO(haya14busa): Remove later when all participants data are in database.
			fillInDummyMembers(false, r.Teams[i])
		}
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
