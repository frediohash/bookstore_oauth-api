package app

import (
	"github.com/frediohash/bookstore_oauth-api/domain/service"
	"github.com/frediohash/bookstore_oauth-api/repository/db"
)

func StartApplication() {
	atService := service.NewService(db.NewRepository())
}
