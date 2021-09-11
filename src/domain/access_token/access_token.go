package access_token

// while accessed by other devices, we need client-id to make expiration of sessions
// Web Frontend - Client-Id : 123
// Android App - Client-Id : 234

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
}
