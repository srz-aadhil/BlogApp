package dto

import "time"

type CreatedUpdatedResponse struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy *int      `json:"created_by"`
	UpdatedBy int       `json:"updated_by"`
}

type DeleteResponse struct {
	DeletedBy int `json:"deleted_by"`
	DeletedAt int `json:"deleted_at"`
}
