package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CustomerID    uint            `json:"customer_id"`
	TotalAmount   float64         `json:"total_amount"`
	Items         []TransactionItem `json:"items" gorm:"foreignKey:TransactionID"`
}

type TransactionItem struct {
	gorm.Model
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
}

type TransactionRequest struct {
	CustomerID uint `json:"customer_id" binding:"required"`
	Items      []TransactionItemRequest `json:"items" binding:"required,min=1"`
}

type TransactionItemRequest struct {
	ProductID uint    `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Price     float64 `json:"price" binding:"required,min=0"`
}

type TransactionResponse struct {
	ID          uint       `json:"id"`
	CustomerID  uint       `json:"customer_id"`
	TotalAmount float64    `json:"total_amount"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
