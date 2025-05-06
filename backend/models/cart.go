package models

import "ecommerce/config"

type CartItem struct {
    ID        int `json:"id" gorm:"primaryKey"`
    UserID    int `json:"user_id"`
    ProductID int `json:"product_id"`
    Quantity  int `json:"quantity"`
}

// Menambahkan item ke keranjang
func AddToCart(userID, productID, quantity int) error {
    var cartItem CartItem
    result := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).First(&cartItem)

    if result.Error != nil {
        // Jika produk belum ada di keranjang, tambahkan baru
        cartItem = CartItem{
            UserID:    userID,
            ProductID: productID,
            Quantity:  quantity,
        }
        return config.DB.Create(&cartItem).Error
    }

    // Jika sudah ada, update jumlahnya
    cartItem.Quantity += quantity
    return config.DB.Save(&cartItem).Error
}

// Mengambil semua item di keranjang user
func GetCartItems(userID int) ([]CartItem, error) {
    var cart []CartItem
    err := config.DB.Where("user_id = ?", userID).Find(&cart).Error
    return cart, err
}
