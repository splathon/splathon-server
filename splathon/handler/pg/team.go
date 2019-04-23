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
		return h.db.Where("event_id = ? AND team_id IS NOT NULL", eventID).Order("id asc").Find(&participants).Error
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	team2members := buildTeam2Members(participants)

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
	var eg errgroup.Group
	var (
		t            Team
		participants []*Participant
	)
	eg.Go(func() error {
		return h.db.Where("id = ?", params.TeamID).Find(&t).Error
	})
	eg.Go(func() error {
		return h.db.Where("team_id = ?", params.TeamID).Order("id asc").Find(&participants).Error
	})
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	team := convertTeam(&t)
	if len(participants) > 0 {
		team.Members = make([]*models.Member, len(participants))
		for i, p := range participants {
			team.Members[i] = convertParticipant2TeamMember(p)
		}
	} else {
		// TODO(haya14busa): Remove later when all participants data are in database.
		fillInDummyMembers(true, team)
	}
	return team, nil
}

func buildTeam2Members(ps []*Participant) map[int64][]*models.Member {
	t2ms := make(map[int64][]*models.Member)
	for _, p := range ps {
		if !p.TeamId.Valid {
			continue
		}
		teamID := p.TeamId.Int64
		if _, ok := t2ms[teamID]; !ok {
			t2ms[teamID] = make([]*models.Member, 0)
		}
		t2ms[teamID] = append(t2ms[teamID], convertParticipant2TeamMember(p))
	}
	return t2ms
}
