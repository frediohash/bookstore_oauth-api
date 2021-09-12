package db

import (
	"github.com/frediohash/bookstore_oauth-api/src/domain/access_token"
	"github.com/frediohash/bookstore_oauth-api/utils/errors"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(id string) (*access_token.AccessToken, *errors.RestErr)
}
type dbRepository struct {
}

func (r *dbRepository) GetById(id string) (*access_token.AccessToken, *errors.RestErr) {
	return nil, nil
}
