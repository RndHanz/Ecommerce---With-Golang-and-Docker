package controllers

import (
	"ecommerce/models"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
    products, err := models.GetAllProducts()
    if err != nil {
        http.Error(w, "Failed to get products", http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}
