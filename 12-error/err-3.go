package main

import (
	"database/sql"
	"errors"
	"fmt"
)

type connection struct {
	*sql.DB
	data string
}

var ErrNoConn = errors.New("can't connect to DB")
var ErrNoRecordFound = errors.New(" record not found")

func (c connection) searchrecord() error {
	db, err := sql.Open("test", "test")
	if err != nil {
		return fmt.Errorf("%v%w", err, ErrNoConn)
	}
	fmt.Println(db)
	//return nil
	_, err = db.Exec("select * from student")
	if err != nil {
		return fmt.Errorf("%v%w", err, ErrNoRecordFound)
	}
	return nil

}
func main() {
	var c connection
	err := c.searchrecord()
	fmt.Println(err)
	if errors.Is(err, ErrNoConn) {
		fmt.Println("please connect to db")
		return
	}
	if errors.Is(err, ErrNoRecordFound) {
		fmt.Println("create a file to search")
		return
	}

}
