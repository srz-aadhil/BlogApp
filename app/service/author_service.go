package service

import (
	"blog/app/dto"
	"blog/app/repo"
	"blog/pkg/e"
	"net/http"
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
	authorRepo repo.AuthorRepo
}

func NewAuthorService(authorRepo repo.AuthorRepo) AuthorService {
	return &AuthorServiceImpl{
		authorRepo: authorRepo,
	}
}

func (s *AuthorServiceImpl) GetAuthor(r *http.Request) (*dto.AuthorResponse, error) {
	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return nil, e.NewError(e.ErrInvalidRequest, "author request parse error", err)
	}

	if err := req.Validate(); err != nil {
		return nil, e.NewError(e.ErrValidateRequest, "author request validation failed", err)
	}

	a, err := s.authorRepo.GetOne(req.ID)
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "not found author with id", err)
	}
	var author dto.AuthorResponse

	author.ID = a.ID
	author.Name = a.Name
	author.CreatedBy = a.CreatedBy
	author.CreatedAt = a.CreatedAt
	author.UpdatedBy = a.UpdatedBy
	author.UpdatedAt = a.UpdatedAt
	author.DeletedBy = a.DeletedBy
	author.DeletedAt = a.DeletedAt

	return &author, nil

}

func (s *AuthorServiceImpl) GetAuthors() (*[]dto.AuthorResponse, error) {
	results, err := s.authorRepo.GetAll()
	if err != nil {
		return nil, e.NewError(e.ErrResourceNotFound, "authors parsing error", err)
	}
	var authorslist []dto.AuthorResponse
	for _, val := range *results {

		var authorResp dto.AuthorResponse
		authorResp.ID = val.ID
		authorResp.Name = val.Name
		authorResp.CreatedAt = val.CreatedAt
		authorResp.CreatedBy = val.CreatedBy
		authorResp.UpdatedAt = val.UpdatedAt
		authorResp.UpdatedBy = val.UpdatedBy
		authorResp.DeletedAt = val.DeletedAt
		authorResp.DeletedBy = val.DeletedBy

		authorslist = append(authorslist, authorResp)
	}
	return &authorslist, nil

}

func (s *AuthorServiceImpl) DeleteAuthor(r *http.Request) error {
	req := &dto.AuthorRequest{}
	if err := req.Parse(r); err != nil {
		return e.NewError(e.ErrInvalidRequest, "author delete parse error", err)
	}

	if err := req.Validate(); err != nil {
		return e.NewError(e.ErrInvalidRequest, "author deletion validate error", err)
	}

	if err := s.authorRepo.Delete(req.ID); err != nil {
		return e.NewError(e.ErrResourceNotFound, "author not found with id", err)
	}
	return nil
}
func (s *AuthorServiceImpl) CreateAuthor(r *http.Request) (int64, error) {
	body := &dto.AuthorCreateRequest{}
	if err := body.Parse(r); err != nil {
		return 0, e.NewError(e.ErrDecodeRequestBody, "author create parse error", err)
	}

	if err := body.Validate(); err != nil {
		return 0, e.NewError(e.ErrValidateRequest, "author create validation error", err)
	}

	authorID, err := s.authorRepo.Create(body)
	if err != nil {
		return 0, e.NewError(e.ErrInvalidRequest, "author creation error", err)
	}
	return authorID, nil

}

func (s *AuthorServiceImpl) UpdateAuthor(r *http.Request) error {
	body := &dto.AuthorUpdateRequest{}
	if err := body.Parse(r); err != nil {
		return e.NewError(e.ErrDecodeRequestBody, "update request decode error", err)
	}

	if err := body.Validate(); err != nil {
		return e.NewError(e.ErrValidateRequest, "author update validation failed", err)
	}

	if err := s.authorRepo.Update(body); err != nil {
		return e.NewError(e.ErrInternalServer, "author updation error", err)
	}
	return nil
}
