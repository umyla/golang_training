package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// const (
// 	host     = "batyr.db.elephantsql.com"
// 	port     = 5432
// 	user     = "jmuuqeva"
// 	password = "yILvHcvbUFs7fUs31XWaFCylEorRC2LH"
// 	dbname   = "jmuuqeva"
// )

// var db *sql.DB

// func main() {
// 	var err error
// 	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
// 		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	db, err = sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()
// 	//Insert(ctx)
// 	//Insert2(ctx)
// 	//querySingleRecords(ctx)
// 	QueryMultipleRecords(ctx)
// }
// func Insert(ctx context.Context) {
// 	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
// 					 VALUES ($1, $2, $3, $4)`

// 	res, err := db.ExecContext(ctx, sqlStatement, 33, "abc@email.com", "dev", "kumar")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(res.RowsAffected())

// }

// func Insert2(ctx context.Context) {
// 	sqlStatement := `INSERT INTO users (age, email, first_name,last_name)
// 					 VALUES ($1, $2, $3, $4)
// 					 RETURNING id,email`
// 	var (
// 		id    int
// 		email string
// 	)
// 	row := db.QueryRowContext(ctx, sqlStatement, 33, email, "abc", "kumar")
// 	err := row.Scan(&id, &email)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(id, email)
// }
// func querySingleRecords(ctx context.Context) {
// 	sqlStatement := `Select id, email FROM users where id = $1;`

// 	var (
// 		id    int
// 		email string
// 	)
// 	err := db.QueryRowContext(ctx, sqlStatement, 30).Scan(&id, &email)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	fmt.Println(id, email)

// }

// func QueryMultipleRecords(ctx context.Context) {
// 	q := `Select id, email FROM users`
// 	rows, err := db.QueryContext(ctx, q)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		var (
// 			id    int
// 			email string
// 		)

// 		err = rows.Scan(&id, &email)
// 		if err != nil {
// 			log.Println(err)
// 		}

// 		fmt.Println(id, email)

// 	}

// }
