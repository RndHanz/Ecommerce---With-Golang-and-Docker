package controllers

import (
	"ecommerce/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func Checkout(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }
    
    userID, _ := strconv.Atoi(r.URL.Query().Get("user_id"))
    total, _ := strconv.ParseFloat(r.URL.Query().Get("total"), 64)
    
    orderID, err := models.CreateOrder(userID, total)
    if err != nil {
        http.Error(w, "Failed to create order", http.StatusInternalServerError)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Order created successfully",
        "order_id": orderID,
    })
}
