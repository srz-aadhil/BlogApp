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

	// User
	userRepo := repo.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	r := chi.NewRouter()
	r.Route("/blogs", func(r chi.Router) {
		r.Post("/create", blogController.CreateBlog)
		r.Get("/", blogController.GetAllBlogs)
		r.Get("/{id}", blogController.GetOneBlog)
		r.Put("/{id}", blogController.UpdateBlog)
		r.Delete("/{id}", blogController.DeleteBlog)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/create", userController.CreateUser)
		r.Get("/", userController.GetAllUsers)
		r.Get("/{id}", userController.GetUser)
		r.Put("/{id}", userController.UpdateUser)
		r.Delete("/{id}", userController.DeleteUser)
	})

	r.Route("/authors", func(r chi.Router) {
		r.Post("/create", authorController.CreateAuthor)
		r.Get("/", authorController.GetaAllAuthors)
		r.Get("/{id}", authorController.GetOneAuthor)
		r.Put("/{id}", authorController.UpdateAuthor)
		r.Delete("/{id}", authorController.DeleteAuthor)

	})
	return r
}
