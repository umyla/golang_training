package database

import "fmt"

type MySqlConn struct {
	db string
}

func NewMySqlConn(conn string) *MySqlConn {
	return &MySqlConn{db: conn}
}

type DbConfig interface {
	Read() string
	Update() string
	Close()
}
type Config struct {
	DB DbConfig
}

func (c Config) Login() {
	fmt.Println("checking for the password from the db")
	fmt.Println(c.DB.Read())
}

func (c *MySqlConn) Read() string {
	return "Reading from Sql"
}
func (c *MySqlConn) Update() string {
	return "updating from Sql"
}
func (c *MySqlConn) Close() {
	c.db = ""
	fmt.Println("closing Sql")
}
