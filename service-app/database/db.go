package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "batyr.db.elephantsql.com"
	port     = 5432
	user     = "jmuuqeva"
	password = "yILvHcvbUFs7fUs31XWaFCylEorRC2LH"
	dbname   = "jmuuqeva"
)

func Open() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
