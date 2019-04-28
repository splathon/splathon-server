package pg

import (
	"context"
	"sort"

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
	// Team ID to Participants.
	tid2ps := make(map[int64][]*Participant)
	for _, p := range ps {
		if !p.TeamId.Valid {
			continue
		}
		teamID := p.TeamId.Int64
		if _, ok := tid2ps[teamID]; !ok {
			tid2ps[teamID] = make([]*Participant, 0)
		}
		tid2ps[teamID] = append(tid2ps[teamID], p)
	}
	// Team ID to Members.
	tid2ms := make(map[int64][]*models.Member)
	for teamID, ps := range tid2ps {
		// Sort participants by order in team.
		sort.Slice(ps, func(i, j int) bool {
			return ps[i].OrderInTeam.Int64 < ps[j].OrderInTeam.Int64
		})
		tid2ms[teamID] = make([]*models.Member, len(ps))
		for i, p := range ps {
			tid2ms[teamID][i] = convertParticipant2TeamMember(p)
		}
	}
	return tid2ms
}
