package database

import "fmt"

type MyPostgresConn struct {
	db string
}

func NewMyPostgresConn(conn string) *MyPostgresConn {
	return &MyPostgresConn{db: conn}
}

func (c *MyPostgresConn) Read() string {
	return "Reading from Postgres"
}
func (c *MyPostgresConn) Update() string {
	return "updating from Postgres"
}
func (c *MyPostgresConn) Close() {
	c.db = ""
	fmt.Println("closing Postgres")
}
