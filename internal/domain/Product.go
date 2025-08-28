package domain

import (
	"time"
)

type ProductInputDTO struct {
	Id          string    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Value       float64   `json:"value" validate:"required"`
}

type ProductOutputDTO struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Value       float64   `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
