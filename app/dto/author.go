package dto

type AuthorResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	CreatedUpdatedResponse
	DeleteResponse
}
