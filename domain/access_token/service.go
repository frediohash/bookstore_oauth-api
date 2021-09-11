package access_token

import (
	"github.com/frediohash/bookstore_users-api/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
type service struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}
func (s *service) GetById(string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}
