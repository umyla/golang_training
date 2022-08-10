package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	var err error
	psqlInfo := "root:Appledad1212#@tcp(127.0.0.1:3306)/users"
	db, err := sql.Open("mysql", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	Insert()
	//Insert2()
	//Update()
	//delete()
	//querySingleRecord()
	//queryMultipleRecords()
}

func Insert() {
	sqlStatement := `Insert INTO users(age,email,first_name,last_name)
	values(? ? ? ?)
	`

	_, err := db.Exec(sqlStatement, 31, "dev@email.com", "dev", "abc")
	if err != nil {
		log.Fatal(err)
	}
}
func Insert2() {
	sqlStatement := `Insert INTO users(age,email,first_name,last_name)
	values(? ? ? ?)`

	res, err := db.Exec(sqlStatement, 32, "diwakar@gmail.com", "Diwakar", "Singh")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.LastInsertId())
}
func Update() {
	sqlStatement := `
	UPDATE users
	SET last_name=?
	where id = ?
	`
	res, err := db.Exec(sqlStatement, "pqk", 2)
	if err != nil {
		log.Fatal(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
func Delete() {
	sqlStatement := `
Delete FROM users
where last_name=?
`
	res, err := db.Exec(sqlStatement, "xyz")
	if err != nil {
		log.Fatal(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted", count)
}
func querySingleRecord() {
	sqlStatement := `Select id,email from user where id=?`
	var (
		id    int
		email string
	)
	row := db.QueryRow(sqlStatement, 2)
	err := row.Scan(&id, &email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("no rows returned")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}
}
func queryMultipleRecords() {
	rows, err := db.Query("Select id,first_name from users LIMIT 3")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id        int
			firstname string
		)
		err = rows.Scan(&id, &firstname)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, firstname)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err)
	}
}
