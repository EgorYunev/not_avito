package data

import (
	"database/sql"
	"log"

	"github.com/EgorYunev/not_avito/config"
	_ "github.com/lib/pq"
)

func Start() *sql.DB {
	db, err := sql.Open("postgres", config.DataBase)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return db
}
