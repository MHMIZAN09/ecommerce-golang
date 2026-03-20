package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.With(
		http.HandlerFunc(handlers.GetProductsHandler),
	))

	mux.Handle("POST /products", manager.With(
		http.HandlerFunc(
			handlers.CreateProductHandler,
		),
	))

	mux.Handle("GET /products/{id}", manager.With(
		http.HandlerFunc(handlers.GetProductByIDHandler),
	))

	mux.Handle("PUT /products/{id}", manager.With(
		http.HandlerFunc(handlers.UpdateProductHandler),
	))

	mux.Handle("DELETE /products/{id}", manager.With(
		http.HandlerFunc(handlers.DeleteProduct),
	))

}
