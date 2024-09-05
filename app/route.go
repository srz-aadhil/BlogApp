package app

import (
	"blog/app/controller"
	"blog/app/repo"
	"blog/app/service"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func apiRouter(db *sql.DB) chi.Router {
	// Author
	authorRepo := repo.NewAuthorRepo(db)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

	// Blog
	blogRepo := repo.NewBlogRepo(db)
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", blogController.GetAllBlogs)
		r.Get("/{id}", blogController.GetOneBlog)
	})
	// userController := controller.NewUserController()

	// r.Route("/users", func(r chi.Router) {
	// 	r.Get("/", userController.GetAllUsers)
	// 	r.Get("/{id}", userController.GetOneUser)
	// })

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", authorController.GetaAllAuthors)
		r.Get("/{id}", authorController.GetOneAuthor)
	})
	return r
}
