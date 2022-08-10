package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // sql.Open using it internally
)

var db *sql.DB //bad idea in production

func main() {
	// %s ->string
	// %d -> int
	var err error
	psqlInfo := "root:Sma121200@tcp(127.0.0.1:3306)/users" // username:pass@tcp(127.0.0.1:3306)/dbname

	db, err = sql.Open("mysql", psqlInfo)
	if err != nil {

		panic(err)
	}
	defer db.Close() // one time only in defer

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	//Insert()
	//Insert2()
	//Update()
	//delete()
	//querySingleRecord()
	queryMultipleRecords()
}

func Insert() {

	sqlStatement := `Insert INTO users (age, email, first_name, last_name)
					Values (? , ? , ? , ?)
`
	_, err := db.Exec(sqlStatement, 31, "dev@email.com", "dev", "xyz")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(res.LastInsertId()) not supported

}

func Insert2() {

	sqlStatement := `
	Insert INTO users (age, email, first_name, last_name)
	Values (? , ?, ?, ? )
	`

	res, err := db.Exec(sqlStatement, 32, "diwakar@gmail.com", "Diwakar", "Singh")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.LastInsertId())

}

func Update() {
	sqlStatement := `
	UPDATE users
	SET last_name = ?
	where id = ?
`
	res, err := db.Exec(sqlStatement, "efgh", 2)
	if err != nil {
		log.Fatal(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
func delete() {

	sqlStatement := `
	Delete FROM users
	where last_name = ?;
`
	res, err := db.Exec(sqlStatement, "xyz")
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted", count)

}

func querySingleRecord() {

	sqlStatement := `Select id,email FROM users where id = ?`
	var (
		id    int
		email string
	)
	row := db.QueryRow(sqlStatement, 2)
	err := row.Scan(&id, &email)
	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows returned")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}
}

func queryMultipleRecords() {
	rows, err := db.Query("SELECT first_name from users LIMIT 3")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() { // return one row if present otherwise stop the loop

		var (
			//id        int
			firstName string
		)
		err = rows.Scan(&firstName) // scanning  data from each row

		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(firstName)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
