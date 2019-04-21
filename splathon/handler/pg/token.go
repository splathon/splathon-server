package pg

import (
	"encoding/base64"
	"encoding/json"
)

type TokenSession struct {
	// Unix timestamp in seconds which this session token is created at.
	CreatedTimestampSec int64 `json:"created_timestamp_sec"`
	IsAdmin             bool  `json:"is_admin,omitempty"`
	// Optional.
	TeamID int64 `json:"team_id,omitempty"`
	// Optional.
	SlackUserID string `json:"slack_userid,omitempty"`
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
}

func NewTokenManager(cipher Cipher) *TokenManager {
	return &TokenManager{cipher: cipher}
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
