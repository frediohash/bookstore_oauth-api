package access_token

import "github.com/frediohash/bookstore_oauth-api/utils/errors"

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
type service struct {
}