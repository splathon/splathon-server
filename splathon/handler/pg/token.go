package pg

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/splathon/splathon-server/splathon/serror"
)

type Env int

const (
	ENV_UNKNOWN Env = iota
	ENV_PROD
	ENV_DEV
)

type TokenSession struct {
	// Unix timestamp in seconds which this session token is created at.
	CreatedTimestampSec int64 `json:"created_timestamp_sec"`
	IsAdmin             bool  `json:"is_admin,omitempty"`
	// Optional.
	TeamID int64 `json:"team_id,omitempty"`
	// Optional.
	SlackUserID string `json:"slack_userid,omitempty"`
	// Optional.
	Env Env `json:"env,omitempty"`
}

// Cipher is crypt interface to encrypt/decrypt cookie.
type Cipher interface {
	Encrypt(plaintext []byte) ([]byte, error)
	Decrypt(ciphertext []byte) ([]byte, error)
}

// TokenManager manages API token.
// Instead of having state and save session data in database, the token itself
// carries necessary token session data for the sake of simplicity of
// maintenance.
type TokenManager struct {
	cipher Cipher
	env    Env
}

func NewTokenManager(cipher Cipher, env Env) *TokenManager {
	return &TokenManager{cipher: cipher, env: env}
}

func (tm *TokenManager) NewToken(isAdmin bool, teamID int64, slackUserID string) TokenSession {
	return TokenSession{
		IsAdmin:             isAdmin,
		TeamID:              teamID,
		SlackUserID:         slackUserID,
		Env:                 tm.env,
		CreatedTimestampSec: time.Now().Unix(),
	}
}

func (tm *TokenManager) Marhal(t TokenSession) (string, error) {
	j, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	crypted, err := tm.cipher.Encrypt(j)
	if err != nil {
		return "", err
	}
	token := base64.StdEncoding.EncodeToString(crypted)
	return token, nil
}

func (tm *TokenManager) Unmarhal(token string) (*TokenSession, error) {
	crypted, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, err
	}
	plaintxt, err := tm.cipher.Decrypt(crypted)
	if err != nil {
		return nil, err
	}
	var t TokenSession
	if err := json.Unmarshal(plaintxt, &t); err != nil {
		return nil, err
	}
	return &t, nil
}

func (tm *TokenManager) ValidateAdminToken(token string) error {
	t, err := tm.Unmarhal(token)
	if err != nil {
		return &serror.Error{Code: http.StatusBadRequest, Message: fmt.Sprintf("invalid token: %v", err)}
	}
	if t.IsAdmin && t.Env == tm.env {
		return nil
	}
	return &serror.Error{Code: http.StatusUnauthorized, Message: "The request user is not authorized."}
}
