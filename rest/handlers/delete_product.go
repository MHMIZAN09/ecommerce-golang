package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // string

	id, err := strconv.Atoi(idStr) // int
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	database.Delete(id)

	utils.SendData(w, "successfully Deleted product", 201)

}
