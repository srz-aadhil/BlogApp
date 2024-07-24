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
	// err := createUser("sukunan@gmail.com", "password123")
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

	// //Create Blog ***********************************************************************************
	// err = createBlog("Blog1", 3, "My first blog", 2, 3)
	// if err != nil {
	// 	fmt.Printf("blog creation failed due to: %s", err)
	// } else {
	// 	fmt.Printf("blog creation successfull")
	// }

	// Delete blog
	// err = deleteBlog(6, 1)
	// if err != nil {
	// 	fmt.Printf("blog deletion failed due to %s:", err)
	// } else {
	// 	fmt.Printf("blog deletion successfull")
	// }

	//Read blog
	// title, content, authorid, created_at, updated_at, err := readBlog(10)
	// if err != nil {
	// 	fmt.Printf("fetching blog failed due to: %s", err)
	// } else {
	// 	fmt.Printf("title: %s \n content: %s\n authorid: %d \n Created at: %s \n Updated at: %s \n", title, content, authorid, created_at, updated_at)
	// }

	//Update blog*********************************************************************************
	// err = updateBlog(10, "updated title", "updated blog content")
	// if err != nil {
	// 	log.Printf("blog updation failed due to : %s", err)
	// } else {
	// 	fmt.Println("successfully updated")
	// }

	//Read All Blogs*****************************************************************************
	blogs, err := readAllBlogs()
	if err != nil {
		log.Printf("getting blogs failed due to : %s", err)
	} else {
		for _, blog := range blogs {
			fmt.Printf("BlogId : %d \n Title: %s \n Content: %s \n AuthorId: %d \n Created at: %s \n Updated at: %s \n", blog.id, blog.title, blog.content, blog.authorid, blog.created_at, blog.updated_at)

		}
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

func readBlog(id uint16) (string, string, uint16, time.Time, time.Time, error) {
	var (
		title      string
		content    string
		authorid   uint16
		created_at time.Time
		updated_at time.Time
	)

	query := `SELECT title,content,author_id,created_at,updated_at FROM blogs WHERE id = $1 AND status=2`

	if err := db.QueryRow(query, id).Scan(&title, &content, &authorid, &created_at, &updated_at); err != nil {
		return "", "", 0, time.Time{}, time.Time{}, err
	}
	return title, content, authorid, created_at, updated_at, nil
}

func updateBlog(id uint16, title string, content string) error {
	query := `UPDATE blogs 
			 SET title= $2,content= $3,updated_at= $4 
			 WHERE id= $1 AND status IN (1,2)`

	result, err := db.Exec(query, id, title, content, time.Now().UTC())
	if err != nil {
		return fmt.Errorf("query execution failed due to : %s", err)
	}
	isAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("no affected rows due to: %s", err)
	}
	if isAffected == 0 {
		return fmt.Errorf("no entry with id %d or status in 1 or 2", id)
	}
	return nil
}

type blog struct {
	id         uint16
	title      string
	authorid   uint16
	content    string
	created_at time.Time
	updated_at time.Time
}

func readAllBlogs() ([]blog, error) {

	query := `SELECT id,title,author_id,content,created_at,updated_at
			  FROM blogs`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed :%d", err)
	}

	defer rows.Close()

	var blogs []blog
	for rows.Next() {
		var blog blog
		if err := rows.Scan(&blog.id, &blog.title, &blog.authorid, &blog.content, &blog.created_at, &blog.updated_at); err != nil {
			return nil, fmt.Errorf("row scan failed due to : %w", err)
		}
		blogs = append(blogs, blog)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration failed due to : %w", err)
	}
	return blogs, nil
}
