package app

import (
	"blog/app/controller"
	"blog/app/repo"
	"blog/app/service"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func apiRouter(db *sql.DB) chi.Router {
	blogController := controller.NewBlogController()
	authorRepo := repo.NewAuthorRepo(db)
	authorService := service.NewAuthorService(authorRepo)
	authorController := controller.NewAuthorController(authorService)

	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Get("/", blogController.GetAllBlogs)
		r.Get("/{id}", blogController.GetOneBlog)
	})
	userController := controller.NewUserController()

	r.Route("/users", func(r chi.Router) {
		r.Get("/", userController.GetAllUsers)
		r.Get("/{id}", userController.GetOneUser)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Get("/", authorController.GetaAllAuthors)
		r.Get("/{id}", authorController.GetOneAuthor)
	})
	return r
}
