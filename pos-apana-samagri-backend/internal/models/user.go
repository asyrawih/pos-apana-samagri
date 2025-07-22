package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole string

const (
	AdminRole   UserRole = "admin"
	CashierRole UserRole = "cashier"
	ManagerRole UserRole = "manager"
)

// String returns the string representation of the role
func (r UserRole) String() string {
	return string(r)
}

// Valid checks if the role is valid
func (r UserRole) Valid() bool {
	switch r {
	case AdminRole, CashierRole, ManagerRole:
		return true
	}
	return false
}

// GormDataType implements the GORM data type interface
func (UserRole) GormDataType() string {
	return "varchar(20)"
}

type User struct {
	gorm.Model
	Email     string   `json:"email" gorm:"uniqueIndex;not null"`
	Password  string   `json:"-" gorm:"not null"`
	FirstName string   `json:"first_name" gorm:"not null"`
	LastName  string   `json:"last_name" gorm:"not null"`
	Role      UserRole `json:"role" gorm:"type:varchar(20);default:'cashier';check:role IN ('admin', 'cashier', 'manager')"`
	IsActive  bool     `json:"is_active" gorm:"default:true"`
}

type RegisterRequest struct {
	Email     string   `json:"email" binding:"required,email"`
	Password  string   `json:"password" binding:"required,min=8"`
	FirstName string   `json:"first_name" binding:"required"`
	LastName  string   `json:"last_name" binding:"required"`
	Role      UserRole `json:"role" binding:"oneof=admin cashier manager"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      UserInfo  `json:"user"`
}

type UserInfo struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Role      UserRole  `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}
