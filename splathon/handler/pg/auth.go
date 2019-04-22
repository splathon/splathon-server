package pg

import (
	"context"
	"errors"
	"net/http"
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

	var p Participant
	if err := h.db.Select("slack_user_id, team_id").Where("slack_username = ? AND raw_password = ?", params.Request.UserID, params.Request.Password).Find(&p).Error; err != nil {
		return nil, errors.New("login failed. ID or Password is wrong. (general login feature has not been implemented yet except admin login)")
	}
	if p.SlackUserId == "" {
		return nil, errors.New("login failed. slack user ID not found.")
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
