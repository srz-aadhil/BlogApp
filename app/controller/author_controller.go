package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"log"
	"net/http"
)

type AuthorController interface {
	CreateAuthor(w http.ResponseWriter, r *http.Request)
	UpdateAuthor(w http.ResponseWriter, r *http.Request)
	GetaAllAuthors(w http.ResponseWriter, r *http.Request)
	GetOneAuthor(w http.ResponseWriter, r *http.Request)
	DeleteAuthor(w http.ResponseWriter, r *http.Request)
}

type AuthorControllerImpl struct {
	authorService service.AuthorService
}

func NewAuthorController(authorService service.AuthorService) AuthorController {
	return &AuthorControllerImpl{
		authorService: authorService,
	}
}

func (c *AuthorControllerImpl) GetaAllAuthors(w http.ResponseWriter, r *http.Request) {
	allAuthors, err := c.authorService.GetAuthors()
	if err != nil {
		log.Printf("can't get authors due to : %s", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, allAuthors)
}

func (c *AuthorControllerImpl) GetOneAuthor(w http.ResponseWriter, r *http.Request) {
	authorResponse, err := c.authorService.GetAuthor(r)
	if err != nil {
		log.Printf("can't get author due to: %s", err)
		api.Fail(w, http.StatusBadRequest, "failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, authorResponse)

}

func (c *AuthorControllerImpl) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	authorID, err := c.authorService.CreateAuthor(r)
	if err != nil {
		log.Printf("author creation failed due to %v", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}

	api.Success(w, http.StatusOK, authorID)
}

func (c *AuthorControllerImpl) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.UpdateAuthor(r); err != nil {
		log.Printf("author updation failed due to %v", err)
		api.Fail(w, http.StatusInternalServerError, "failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Author Updation Success")
}

func (c *AuthorControllerImpl) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	if err := c.authorService.DeleteAuthor(r); err != nil {
		log.Printf("author deletion failed due to %v", err)
		api.Fail(w, http.StatusInternalServerError, "Failed", err.Error())
		return
	}
	api.Success(w, http.StatusOK, "Author deletion successfull")
}
