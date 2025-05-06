package main

import (
	"ecommerce/config"
	"ecommerce/routes"
	"fmt"
	"net/http"
)

func main() {
    config.ConnectDB() // Pastikan ini sesuai dengan database kamu
    routes.RegisterRoutes()

    fmt.Println("Server running on port 8080...")
    http.ListenAndServe(":8080", nil)
}
