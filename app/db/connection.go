package db

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	user     = "postgres"
	password = "password"
	host     = "localhost"
	port     = 5432
	dbname   = "blogapp"
)

func IniDb() (*sql.DB, error) {
	connectionString := fmt.Sprintf("user=%v password=%v host=%v port=%d dbname=%v sslmode=disable", user, password, host, port, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection succesfully created")
	return db, nil
}
