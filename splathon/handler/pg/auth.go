package pg

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/splathon/serror"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) Login(ctx context.Context, params operations.LoginParams) (*models.LoginResponse, error) {
	eventID, err := h.queryInternalEventID(params.EventID)
	if err != nil {
		return nil, err
	}

	if h.isAdminLoginReq(params) {
		apiToken, err := h.tm.Marhal(h.tm.NewToken(true, 0, ""))
		if err != nil {
			return nil, err
		}
		return &models.LoginResponse{
			IsAdmin: swag.Bool(true),
			Token:   swag.String(apiToken),
		}, nil
	}

	resp := &models.LoginResponse{
		IsAdmin: swag.Bool(false),
	}

	// Fetch multiple participants as multiple participants are associated with single user id (and associated with the same password).
	// Assuming they are not associated with different teams, use team_id from one of them as a primary team id.
	var ps []*Participant
	if err := h.db.Select("slack_user_id, team_id").Where("event_id = ? AND slack_username = ? AND raw_password = ?", eventID, strings.Trim(*params.Request.UserID, " "), params.Request.Password).Find(&ps).Error; err != nil || len(ps) == 0 {
		return nil, &serror.Error{Code: http.StatusBadRequest, Message: "login failed. ID or Password is wrong"}
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].TeamId.Int64 > ps[j].TeamId.Int64
	})
	p := ps[0]
	if p.SlackUserId == "" {
		return nil, errors.New("login failed. slack user ID not found")
	}

	fmt.Printf("[INFO] Login: %s\n", p.SlackUserId)

	if p.TeamId.Valid {
		var team Team
		if err := h.db.Where("id = ?", p.TeamId.Int64).Find(&team).Error; err != nil {
			return nil, err
		}
		resp.Team = convertTeam(&team)
	}
	apiToken, err := h.tm.Marhal(h.tm.NewToken(false, p.TeamId.Int64, p.SlackUserId))
	if err != nil {
		return nil, err
	}
	resp.Token = swag.String(apiToken)
	return resp, nil
}

func (h *Handler) isAdminLoginReq(params operations.LoginParams) bool {
	return *params.Request.UserID == h.adminID && *params.Request.Password == h.adminPW
}

func (h *Handler) getTokenSession(token string) (*TokenSession, error) {
	t, err := h.tm.Unmarhal(token)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	return t, nil
}

func (h *Handler) checkAdminAuth(token string) error {
	return h.tm.ValidateAdminToken(token)
}
