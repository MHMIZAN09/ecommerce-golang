package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // string

	id, err := strconv.Atoi(idStr) // int
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	for _, product := range database.Products {
		// log.Println(idx, product)

		if product.ID == id {
			utils.SendData(w, product, 200)
			return
		}
	}

	http.Error(w, "Product not found", 404)
}
