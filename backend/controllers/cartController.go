package controllers

import (
	"ecommerce/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func AddToCart(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    
    userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
    productID, _ := strconv.Atoi(r.URL.Query().Get("product_id"))
    quantity, _ := strconv.Atoi(r.URL.Query().Get("quantity"))
    
    err := models.AddToCart(userID, productID, quantity)
    if err != nil {
        http.Error(w, "Failed to add to cart", http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Added to cart successfully"))
}

func GetCart(w http.ResponseWriter, r *http.Request) {
    userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
    
    cart, err := models.GetCartItems(userID)
    if err != nil {
        http.Error(w, "Failed to fetch cart", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(cart)
}
