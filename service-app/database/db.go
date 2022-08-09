package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "diwakar"
	password = "root"
	dbname   = "postgres"
)

func Open() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+"password=%s dbname=%s sslmode=disabled", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	return db, err
}
