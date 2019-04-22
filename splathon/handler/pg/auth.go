package pg

import (
	"context"
	"errors"
	"net/http"
	"sort"
	"time"

	"github.com/go-openapi/swag"
	"github.com/splathon/splathon-server/splathon/serror"
	"github.com/splathon/splathon-server/swagger/models"
	"github.com/splathon/splathon-server/swagger/restapi/operations"
)

func (h *Handler) Login(ctx context.Context, params operations.LoginParams) (*models.LoginResponse, error) {
	token := TokenSession{
		CreatedTimestampSec: time.Now().Unix(),
	}
	if h.isAdminLoginReq(params) {
		token.IsAdmin = true
		apiToken, err := h.tm.Marhal(token)
		if err != nil {
			return nil, err
		}
		return &models.LoginResponse{
			IsAdmin: true,
			Token:   swag.String(apiToken),
		}, nil
	}

	// Fetch multiple participants as multiple participants are associated with single user id (and associated with the same password).
	// Assuming they are not associated with different teams, use team_id from one of them as a primary team id.
	var ps []*Participant
	if err := h.db.Select("slack_user_id, team_id").Where("slack_username = ? AND raw_password = ?", params.Request.UserID, params.Request.Password).Find(&ps).Error; err != nil || len(ps) == 0 {
		return nil, errors.New("login failed. ID or Password is wrong")
	}
	sort.Slice(ps, func(i, j int) bool {
		return ps[i].TeamId.Int64 > ps[j].TeamId.Int64
	})
	p := ps[0]
	if p.SlackUserId == "" {
		return nil, errors.New("login failed. slack user ID not found")
	}
	token.SlackUserID = p.SlackUserId
	if p.TeamId.Valid {
		token.TeamID = p.TeamId.Int64
	}
	apiToken, err := h.tm.Marhal(token)
	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{Token: swag.String(apiToken)}, nil
}

func (h *Handler) isAdminLoginReq(params operations.LoginParams) bool {
	return *params.Request.UserID == h.adminID && *params.Request.Password == h.adminPW
}

func (h *Handler) getTokenSession(token string) (*TokenSession, error) {
	return h.tm.Unmarhal(token)
}

func (h *Handler) checkAdminAuth(token string) error {
	t, err := h.getTokenSession(token)
	if err != nil {
		return err
	}
	if t.IsAdmin {
		return nil
	}
	return &serror.Error{Code: http.StatusUnauthorized, Message: "The request user has no access for the requested operation."}
}
