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

	product := database.Get(id)
	if product == nil {
		utils.SendError(w, "Product not found", 404)
		return
	}


	utils.SendData(w, product, 200)
}
