package models

import "ecommerce/config"

type Product struct {
    ID          int     `json:"id" gorm:"primaryKey"`
    Name        string  `json:"name"`
    Description string  `json:"description"`
    Price       float64 `json:"price"`
    Image       string  `json:"image"`
}

// Mengambil semua produk
func GetAllProducts() ([]Product, error) {
    var products []Product
    err := config.DB.Find(&products).Error
    return products, err
}
