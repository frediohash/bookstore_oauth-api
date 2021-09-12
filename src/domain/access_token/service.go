package access_token

import (
	"github.com/frediohash/bookstore_oauth-api/src/repository/db"
	"github.com/frediohash/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
type service struct {
	repository db.DbRepository
}

func NewService(repository db.DbRepository) Service {
	return &service{
		repository: repository,
	}
}
func (s *service) GetById(id string) (*AccessToken, *errors.RestErr) {
	return nil, nil
}
