package service

import (
	"github.com/frediohash/bookstore_oauth-api/domain/access_token"
	"github.com/frediohash/bookstore_oauth-api/repository/db"
	"github.com/frediohash/bookstore_oauth-api/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
}
type service struct {
	repository db.DbRepository
}

func NewService(repository db.DbRepository) Service {
	return &service{
		repository: repository,
	}
}
func (s *service) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
