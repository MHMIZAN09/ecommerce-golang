package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Please give me post method", 400)
		return
	}
	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		http.Error(w, "Invalid request body", 400)
		return
	}
	newProduct.ID = len(database.Products) + 1
	database.Products = append(database.Products, newProduct)

	utils.SendData(w, newProduct, 201)
}
