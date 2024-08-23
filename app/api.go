package app

import (
	"blog/pkg/api"
	"database/sql"
)

func Start(db *sql.DB) {
	r := apiRouter(db)
	api.StartServer(":8080", r)

}
