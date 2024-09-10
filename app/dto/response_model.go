package dto

import "time"

type CreatedUpdatedResponse struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedBy *int       `json:"created_by,omitempty"`
	UpdatedBy *int       `json:"updated_by,omitempty"`
}

type DeleteResponse struct {
	DeletedBy *int       `json:"deleted_by,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
