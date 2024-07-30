package repo

import "time"

type Model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	CreatedBy uint16
	UpdatedBy uint16
}

type DeleteInfo struct {
	DeletedAt time.Time
	DeletedBy time.Time
}
