package app

import (
	"github.com/frediohash/bookstore_oauth-api/src/domain/access_token"
	"github.com/frediohash/bookstore_oauth-api/src/repository/db"
)

func StartApplication() {
	dbRepo := db.NewRepository()
	atService := access_token.NewService(dbRepo)
}
