package models

import (
	"ecommerce/config"
	"time"
)

type Order struct {
    ID        int       `json:"id" gorm:"primaryKey"`
    UserID    int       `json:"user_id"`
    Total     float64   `json:"total"`
    CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

// Membuat order baru
func CreateOrder(userID int, total float64) (int, error) {
    order := Order{
        UserID: userID,
        Total:  total,
    }
    result := config.DB.Create(&order)
    return order.ID, result.Error
}
