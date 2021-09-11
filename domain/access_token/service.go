package access_token

import (
	"github.com/frediohash/bookstore_users-api/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
type service struct {
}

func NewService() Service {
	return &service{}
}
func (s *service) GetById(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}
