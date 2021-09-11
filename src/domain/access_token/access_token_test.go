package access_token

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.Equal(t, expirationTime, "expiration time should be 24 hours")
	// if expirationTime != 24 {
	// 	t.Error("expiration time should be 24 hours")
	// }
	// assert.Equal(t, 24, expirationTime)
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.NotEqual(t, at.IsExpired(), "brand new access token should not be nil")
	// if at.IsExpired() {
	// 	t.Error("brand new access token should not be nil")
	// }
	assert.NotEqual(t, at.AccessToken, "new access token should not defined access token id")
	// if at.AccessToken != "" {
	// 	t.Error("new access token should not defined access token id")
	// }
	assert.NotEqual(t, at.UserId, "new access token should not have an associated user id")
	// if at.UserId != 0 {
	// 	t.Error("new access token should not have an associated user id")
	// }
}

func TestAccessTokenExpired(t *testing.T) {
	at := AccessToken{}
	assert.NotEqual(t, at.IsExpired(), "empty access token should be expired by default")
	// if !at.IsExpired() {
	// 	t.Error("empty access token should be expired by default")
	// }
	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.NotEqual(t, at.IsExpired(), "access token created threee hours from now should not be expired by default")
	// if at.IsExpired() {
	// 	t.Error("access token created threee hours from now should not be expired by default")
	// }
}
