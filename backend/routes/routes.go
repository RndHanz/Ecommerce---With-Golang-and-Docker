package routes

import (
	"ecommerce/controllers"
	"net/http"
)

func RegisterRoutes() {
    http.HandleFunc("/api/products", controllers.GetProducts)
    http.HandleFunc("/api/cart/add", controllers.AddToCart)
    http.HandleFunc("/api/cart", controllers.GetCart)
    http.HandleFunc("/api/checkout", controllers.Checkout)
}
