package biz

import (
	"time"
)

type Employee struct {
	ID    int64
	Name  string
	Email string
	Role  string

	CreatedAt time.Time
	UpdatedAt time.Time
}
