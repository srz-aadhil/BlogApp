package dto

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	validator "github.com/go-playground/validator/v10"
)

type AuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CreatedUpdatedResponse
	DeleteResponse
}

type AuthorRequest struct {
	ID int `validate:"required"`
}

// for Path param
func (a *AuthorRequest) Parse(r *http.Request) error {
	//get author ID from request
	strID := chi.URLParam(r, "id")
	// converting string ID to int ID
	intID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = intID
	return nil
}

func (a *AuthorRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}

// For Body param
type AuthorCreateRequest struct {
	Name      string `json:"name"`
	CreatedBy int    `json:"created_by"`
}

func (a *AuthorCreateRequest) Parse(r *http.Request) error {
	
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil
}

func (a *AuthorCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil
}

type AuthorUpdateRequest struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UpdatedBy int    `json:"updated_by"`
}

func (a *AuthorUpdateRequest) Parse(r *http.Request) error {
	//Get ID from request
	strID := chi.URLParam(r, "id")
	IntID, err := strconv.Atoi(strID)
	if err != nil {
		return err
	}
	a.ID = IntID
	//Decode to AuthorUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		return err
	}
	return nil

}

func (a *AuthorUpdateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		return err
	}
	return nil

}
