package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type AuthorService interface {
	GetAuthor(r *http.Request) (*dto.AuthorResponse, error)
	GetAuthors() (*[]dto.AuthorResponse, error)
	DeleteAuthor(r *http.Request) error
	CreateAuthor(r *http.Request) (int64, error)
	UpdateAuthor(r *http.Request) error
}

var _ AuthorService = (*AuthorServiceImpl)(nil)

type AuthorServiceImpl struct {
	authorRepo repo.Repo
}

func NewAuthorService(authorRepo repo.Repo) AuthorService {
	return &AuthorServiceImpl{
		authorRepo: authorRepo,
	}
}

func (s *AuthorServiceImpl) GetAuthor(r *http.Request) (*dto.AuthorResponse, error) {

	//get author ID from request
	strID := chi.URLParam(r, "id")
	// converting string ID to int ID
	intID, err := strconv.Atoi(strID)
	//fmt.Printf("author id is :: %d", intID)
	if err != nil {
		return nil, err
	}
	result, err := s.authorRepo.GetOne(intID)
	if err != nil {
		return nil, err
	}

	a, ok := result.(repo.Author)
	if !ok {
		return nil, err

	}

	//fmt.Println("assertion before: ",a)
	var authors dto.AuthorResponse
	authors.ID = a.ID
	authors.Name = a.Name
	authors.CreatedBy = a.CreatedBy
	authors.CreatedAt = a.CreatedAt
	//fmt.Println("::::::" , authres)
	return &authors, nil

}

func (s *AuthorServiceImpl) GetAuthors() (*[]dto.AuthorResponse, error) {
	results, err := s.authorRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var authors []dto.AuthorResponse
	for _, val := range results {
		a, ok := val.(repo.Author)
		if !ok {
			return nil, err
		}

		var author dto.AuthorResponse
		author.ID = a.ID
		author.Name = a.Name
		author.CreatedAt = a.CreatedAt
		author.CreatedBy = a.CreatedBy

		authors = append(authors, author)
	}
	return &authors, nil
}

func (s *AuthorServiceImpl) DeleteAuthor(r *http.Request) error {
	req
}
