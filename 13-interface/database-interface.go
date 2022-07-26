package main

import "learn-go/database"

func main() {
	m := database.NewMySqlConn("localhost")
	c := database.Config{DB: m}
	c.Login()
	p := database.NewMyPostgresConn("localhost")
	c = database.Config{DB: p}
	c.Login()
}
