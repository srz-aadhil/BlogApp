package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgress"
	password = "password"
	host     = "localhost"
	port     = 5432
	dbname   = "BlogDB"
)

var DB *sql.DB
var err error

func main() {
	ConnectionString := fmt.Sprintf("user = %s password = %s host = %s port = %d dbname = %s", user, password, host, port, dbname)

	DB, err = sql.Open("postgres", ConnectionString)

	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	if err := DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Successfully Connected")

}
