package main

import (
	"database/sql"
	"fmt"
	"log"
)

type config struct {
	DbConn string
	data   string
}

var Db string

func NewConfig(DbConn string) (config, error) {
	if DbConn == "" {
		return config{}, sql.ErrConnDone
	}
	c := config{
		DbConn: DbConn,
	}
	return c, nil
}
func (x config) read() {
	if x.DbConn == "" {
		log.Println("not allowed to read")
		return
	}
}
func (x *config) update(data string) {
	if x.DbConn == "" {
		log.Println("not allowed to read")
		return
	}
}
func main() {
	conn, err := NewConfig("localhost:mysql")
	if err != nil {
		log.Println(err)
	}

	conn.update("new data")
	conn.read()
	fmt.Println("connected to db", conn)
}
