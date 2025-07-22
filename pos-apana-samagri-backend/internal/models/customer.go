package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	FirstName string    `json:"first_name" gorm:"not null"`
	LastName  string    `json:"last_name" gorm:"not null"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
}

type CustomerRequest struct {
	FirstName string    `json:"first_name" binding:"required"`
	LastName  string    `json:"last_name" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
}

type CustomerResponse struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
