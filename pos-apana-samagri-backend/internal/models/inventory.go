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

// String returns the string representation of the status
func (s InventoryStatus) String() string {
	return string(s)
}

// Valid checks if the status is valid
func (s InventoryStatus) Valid() bool {
	switch s {
	case InStock, LowStock, OutOfStock:
		return true
	}
	return false
}

// GormDataType implements the GORM data type interface
func (InventoryStatus) GormDataType() string {
	return "varchar(20)"
}

type Inventory struct {
	gorm.Model
	ProductID       uint            `json:"product_id" gorm:"uniqueIndex"`
	Quantity        int             `json:"quantity" gorm:"default:0"`
	Status          InventoryStatus `json:"status" gorm:"type:varchar(20);default:'out_of_stock';check:status IN ('in_stock', 'low_stock', 'out_of_stock')"`
	Location        string          `json:"location"`
	LastStockUpdate time.Time       `json:"last_stock_update"`
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
