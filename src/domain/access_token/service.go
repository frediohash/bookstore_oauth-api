package access_token

import (
	"github.com/frediohash/bookstore_users-api/utils/errors"
	"github.com/micro/go-micro/errors"
)

type Service interface {
	GetById(string) (*AccessToken, *errors.RestErr)
}
