package controller

import (
	"blog/app/service"
	"blog/pkg/api"
	"log"
	"net/http"
)

type AuthorController interface {
	GetaAllAuthors(w http.ResponseWriter, r *http.Request)
	GetOneAuthor(w http.ResponseWriter, r *http.Request)
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

	api.Send(w, http.StatusOK, allAuthors)
}

func (c *AuthorControllerImpl) GetOneAuthor(w http.ResponseWriter, r *http.Request) {
	authorResponse, err := c.authorService.GetAuthor(r)
	if err != nil {
		log.Printf("can't get author due to: %s", err)
		api.Fail(w, http.StatusBadRequest, "failed", err.Error())
		return
	}

	api.Send(w, http.StatusOK, authorResponse)

}
