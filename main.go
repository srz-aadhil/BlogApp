package main

import (
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	password = "srzaadhil"
	host     = "localhost"
	port     = 5432
	dbname   = "blogapp"
)

var db *sql.DB
var err error

func main() {
	connectionString := fmt.Sprintf("user = %s password = %s host = %s port = %d dbname = %s sslmode = disable", user, password, host, port, dbname)

	db, err = sql.Open("postgres", connectionString)
	// DSN parse error or initialisation error
	if err != nil {
		log.Fatal(err)
	}
	// close db connection before main function exit
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database Successfully Connected")

	// User Creation *********************************************************************************
	// Try with UNIQUE usernames..
	// err := createUser("aadhil696@gmail.com", "password007")
	// if err != nil {
	// 	log.Printf("user creation failed due to: %s", err)
	// } else {
	// 	fmt.Printf("User created successfully")
	// }

	// Create Author *********************************************************************************
	// err = createAuthor("author1111")
	// if err != nil {
	// 	fmt.Printf("author creation failed due to: %s", err)
	// } else {
	// 	fmt.Printf("Author created successfully")
	// }

	// Create Blog ***********************************************************************************
	err = createBlog("Blog1", 5, "My first blog", 1, 1)
	if err != nil {
		fmt.Printf("blog creation failed due to: %s", err)
	} else {
		fmt.Printf("blog creation successfull")
	}

	// Delete blog
	err = deleteBlog(6, 1)
	if err != nil {
		fmt.Printf("blog deletion failed due to %s:", err)
	} else {
		fmt.Printf("blog deletion successfull")
	}

}
func createUser(username, password string) error {
	// Generate salt
	salt, err := generateSalt(10)
	if err != nil {
		return fmt.Errorf("error generating salt: %v", err)
	}

	// Password Hashing
	hashedPassword := hashPassword(password, salt)

	query := `INSERT INTO users(username,password,salt)
			  VALUES ($1,$2,$3)`

	_, err = db.Exec(query, username, hashedPassword, salt)
	if err != nil {
		return fmt.Errorf("execution error due to : %s", err)
	}

	return nil
}

// Generate a random salt of the given length
func generateSalt(length uint8) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	saltString := hex.EncodeToString(bytes)
	return saltString, nil
}

// Hashes the password with given salt (SHA-256)
func hashPassword(password, salt string) string {
	hash := sha256.New()
	hash.Write([]byte(password + salt))
	hashBytes := hash.Sum(nil)
	hashedPass := hex.EncodeToString(hashBytes)
	return hashedPass
}

func createAuthor(name string) error {

	query := `INSERT INTO authors(name)
			  VALUES ($1)`

	//if err := db.QueryRow(query, name).Scan(&authorId); err != nil {
	//return 0, err
	//}
	_, err = db.Exec(query, name)
	if err != nil {
		return fmt.Errorf("execution error due to : %s ", err)
	}
	return nil
}

func createBlog(title string, authorId uint16, content string, status int, userId int) error {
	query := `INSERT INTO blogs(title,author_id,content,status,created_by)
			  VALUES($1,$2,$3,$4,$5)`

	_, err = db.Exec(query, title, authorId, content, status, userId)
	if err != nil {
		return fmt.Errorf("query execution failed due to: %s", err)
	}
	return nil
}

func deleteBlog(id, userId uint16) error {
	query := `UPDATE blogs
			  SET deleted_by=$1,deleted_at=$2,status=$3
			  WHERE id=$4`

	_, err = db.Exec(query, userId, time.Now().UTC(), 3, id)
	if err != nil {
		return fmt.Errorf("delete query execution failed due to: %s", err)
	}
	return nil
}
