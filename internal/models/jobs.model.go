package models

import (
	"time"

	"github.com/google/uuid"
)

type Job struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CompanyID   uuid.UUID `json:"company_id" db:"company_id"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

type JobCompany struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CompanyID   uuid.UUID `json:"company_id" db:"company_id"`
	CompanyName string    `json:"company_name" db:"company_name"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
