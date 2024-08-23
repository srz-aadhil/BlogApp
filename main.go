package main

import (
	"blog/app"
	"blog/app/db"
	"log"

	// "log"
	_ "github.com/lib/pq"
)

// const (
// 	user     = "postgres"
// 	password = "srzaadhil"
// 	host     = "localhost"
// 	port     = 5432
// 	dbname   = "blogapp"
// )

// var db *sql.DB
// var err error

func main() {
	db, err := db.IniDb()
	if err != nil {
		log.Fatal(err)
	}
	app.Start(db)
	defer db.Close()

	// Get All blogs
	// r.Get("/blogs", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Got All blogs"))
	// })

	// Get one blog
	// r.Get("/blog/{blog_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Got the blog with id"))
	// })

	// // Update a blog
	// r.Put("/blog/{blog_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Updated blog with id"))
	// })

	// // Delete a blog
	// r.Delete("/blog/{blog_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Deleted blog with id "))
	// })

	// // Create a blog
	// r.Post("/blog", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Created a blog with id"))
	// })

	// // Get All Users
	// r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Got all users"))
	// })

	// // Get One User
	// r.Get("/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Got user with id"))
	// })

	// // Update a user
	// r.Put("/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("User updated with user id-"))
	// })

	// // Delete a user
	// r.Delete("/users/{user_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("User deleted with id-"))
	// })

	// // Create a user
	// r.Post("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Created user with id"))
	// })

	// // Get all authors
	// r.Get("/authors", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Got all users"))
	// })

	// // Get one author
	// r.Get("/authors/{author_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Got one author"))
	// })

	// // Update an author
	// r.Put("/authors/{author_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Updated a author with id-"))
	// })

	// // Delete an author
	// r.Delete("/authors/{author_id}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Deleted an author with id-"))
	// })

	// // Create an author
	// r.Post("/authors", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Created author with id-"))
	// })

	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	// })

	// log.Fatal(http.ListenAndServe(":8080", r))

	// connectionString := fmt.Sprintf("user = %s password = %s host = %s port = %d dbname = %s sslmode = disable", user, password, host, port, dbname)

	// db, err = sql.Open("postgres", connectionString)
	// // DSN parse error or initialisation error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // close db connection before main function exit
	// defer db.Close()

	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Database Successfully Connected")

	// var user repo.User // creating an instance of user
	// var author repo.Author // creating an instance of author
	// var blog repo.Blog //creating an instance of blog

	// User creation
	// 	user.UserName = "django"
	// 	user.Password = "django123"

	// 	userID, err := user.Create(db)
	// 	if err != nil {
	// 		log.Printf("user creation failed due to : %s", err)
	// 	} else {
	// 		fmt.Printf("user created with ID : %d", userID)
	// 	}

	// 	user.ID = 7								// User deletion
	// 	if err := user.Delete(db); err != nil {
	// 		log.Printf("user deletion failed due to :%s", err)
	// 	} else {
	// 		fmt.Printf("user deletion successfull")
	// 	}

	// 	user.UserName = "djangozz"
	// 	user.Password = "django007"
	// 	user.ID = 9
	// 	if err := user.Update(db); err != nil {
	// 		log.Printf("user updation failed due to $: %s", err)
	// 	} else {
	// 		fmt.Printf("user details updated successfully")
	// 	}

	// 	userResult, err := user.GetOne(db) // Get single user details
	// 	if err != nil {
	// 		log.Printf("fetching user details failed due : %s", err)
	// 	} else {
	// 		// Type assertion to convert userResult to Users
	// 		u, ok := userResult.(repo.User)
	// 		if !ok {
	// 			log.Printf("type convertion failed ")
	// 		}
	// 		fmt.Printf("Username : %s \n UserID : %d", u.UserName, u.ID)
	// 	}

	// 	Allusers, err := user.GetAll(db) // Get All users details
	// 	if err != nil {
	// 		log.Printf("all users fetching failed due to :%s", err)
	// 	} else {
	// 		for _, userslist := range Allusers {
	// 			//type assertion to convert userslist to Users type
	// 			au, ok := userslist.(repo.User)
	// 			if !ok {
	// 				log.Printf("type assertion failed")
	// 			} else {
	// 				fmt.Printf(" Username : %s UserID : %d CreatedAt : %v \n", au.UserName, au.ID, au.CreatedAt)
	// 			}
	// 		}
	// 	}

	// Author creation
	// author.Name = "authornew"
	// author.CreatedBy = 7

	// authorID, err := author.Create(db)
	// if err != nil {
	// 	log.Printf("author creation failed due to : %s", err)
	// } else {
	// 	fmt.Printf("Author created with authorID : %d", authorID)
	// }

	// author.Name = "iauthor"
	// author.ID = 6
	// author.UpdatedBy = 7

	// 	if err = author.Update(db); err != nil {
	// 		log.Printf("Author updation failed due to : %s", err)
	// 	} else {
	// 		fmt.Printf("Author updation successfull")
	// 	}

	// 	author.ID = 5
	// 	author.DeletedBy = 7
	// 	if err = author.Delete(db); err != nil {
	// 		log.Printf("author deletion failed due to :%s", err)
	// 	} else {
	// 		fmt.Printf("author deletion successfull")
	// 	}

	// Get single authors

	// singleAuthor, err := author.GetOne(db)
	// 	if err != nil {
	// 		log.Printf("author fetching failed due to : %s", err)
	// 	} else {
	// 		//type assertion of singleAuthor to type
	// 		oneauthor, ok := singleAuthor.(repo.Author)
	// 		if !ok {
	// 			fmt.Printf("type assertion failed")
	// 		} else {
	// 			fmt.Printf(" AuthorName : %s AuthorID : %d ", oneauthor.Name, oneauthor.ID)
	// 		}
	// 	}

	// Get All Authors
	// allAuthors, err := author.GetAll(db)
	// if err != nil {
	// 	log.Printf("author fetching failed due to :%s", err)
	// } else {
	// 	for _, authorslist := range allAuthors {
	// 		//type assertion of allAuthors(interface) to author type
	// 		aa, ok := authorslist.(repo.Author)
	// 		if !ok {
	// 			fmt.Printf("Type assertion failed")
	// 		} else {
	// 			fmt.Printf("Author Name: %s Author ID: %d Created At: %v Created By: %d \n ", aa.Name, aa.ID, aa.CreatedAt, aa.CreatedBy)
	// 		}
	// 	}
	// }

}
