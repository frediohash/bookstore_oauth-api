package app

import (
	"github.com/frediohash/bookstore_oauth-api/repository/db"
)

func StartApplication() {
	atService := ser(db.NewRepository())
}
