package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProductsHandler))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProductHandler)) // route

	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByIDHandler))

	globalRouter := global_router.GlobalRouter(mux)
	fmt.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe(":8080", globalRouter)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
