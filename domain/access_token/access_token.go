package access_token

import (
	"time"

	"github.com/frediohash/bookstore_users-api/utils/errors"
)

// while accessed by other devices, we need client-id to make expiration of sessions
// Web Frontend - Client-Id : 123
// Android App - Client-Id : 234

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

type Repository interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	return now.After(expirationTime)
}
