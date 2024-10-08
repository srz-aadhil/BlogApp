package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type BlogResponse struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	AuthorID int    `json:"authorid"`
	Status   int    `json:"status"`
	CreatedUpdatedResponse
	DeleteResponse
}

// for path param
type BlogRequest struct {
	ID int `validation:"required"`
}

func (b *BlogRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	b.ID = intID
	return nil
}

func (b *BlogRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

type BlogDeleteRequest struct {
	ID        int `json:"id"`
	DeletedBy int `json:"deleted_by"`
}

func (b *BlogDeleteRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	b.ID = intID
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BlogDeleteRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

// for body param
type BlogCreateRequest struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	AuthorID  int    `json:"author_id" validate:"required"` //Author.ID
	Status    int    `json:"status"`
	CreatedBy int    `json:"created_by" validate:"required"` // User.ID
}

func (b *BlogCreateRequest) Parse(r *http.Request) error {
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BlogCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}

type BlogUpdateRequest struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    int    `validate:"required"`
	UpdatedBy int    `json:"updated_by" validate:"required"`
}

func (b *BlogUpdateRequest) Parse(r *http.Request) error {
	strID := chi.URLParam(r, "id")
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}

	b.ID = intID
	if err := json.NewDecoder(r.Body).Decode(b); err != nil {
		return err
	}
	return nil
}

func (b *BlogUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(b); err != nil {
		return err
	}
	return nil
}
