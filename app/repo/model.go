package repo

import "time"

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy *int
	UpdatedBy int
}

type DeleteInfo struct {
	DeletedAt time.Time
	DeletedBy int
}
