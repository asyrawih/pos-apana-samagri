package models

import (
	"time"

	"gorm.io/gorm"
)

type InventoryStatus string

const (
	InStock    InventoryStatus = "in_stock"
	LowStock   InventoryStatus = "low_stock"
	OutOfStock InventoryStatus = "out_of_stock"
)

type Inventory struct {
	gorm.Model
	ProductID uint            `json:"product_id" gorm:"uniqueIndex"`
	Quantity  int             `json:"quantity" gorm:"default:0"`
	Status    InventoryStatus `json:"status" gorm:"type:inventory_status;default:'out_of_stock'"`
	Location  string          `json:"location"`
	LastStockUpdate time.Time  `json:"last_stock_update"`
}

type InventoryUpdateRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}

type InventoryResponse struct {
	ID              uint            `json:"id"`
	ProductID       uint            `json:"product_id"`
	ProductName     string          `json:"product_name,omitempty"`
	Quantity        int             `json:"quantity"`
	Status          InventoryStatus `json:"status"`
	Location        string          `json:"location"`
	LastStockUpdate time.Time       `json:"last_stock_update"`
}

type InventoryAlert struct {
	ID          uint      `json:"id"`
	ProductID   uint      `json:"product_id"`
	ProductName string    `json:"product_name"`
	Message     string    `json:"message"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
