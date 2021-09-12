package app

import (
	"github.com/frediohash/bookstore_oauth-api/domain/access_token"
	"github.com/frediohash/bookstore_oauth-api/repository/db"
)

func StartApplication() {
	dbRepository := db.NewRepository()
	atService := access_token.NewService(dbRepository)
}
