package main

import (
	"database/sql"
	"fmt"
	"log"

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
	// 	blogs, err := GetAll()
	// 	if err != nil {
	// 		log.Printf("getting blogs failed due to : %s", err)
	// 	} else {
	// 		for _, blog := range blogs {
	// 			fmt.Printf("BlogId : %d \n Title: %s \n Content: %s \n AuthorId: %d \n Created at: %s \n Updated at: %s \n", blog.id, blog.title, blog.content, blog.authorid, blog.created_at, blog.updated_at)

	// 		}
	// 	}
}
