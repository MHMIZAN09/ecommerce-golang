package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id") // string

	id, err := strconv.Atoi(idStr) // int
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me valid json", 400)
		return
	}

	newProduct.ID = id
	database.Update(newProduct)

	utils.SendData(w, "successfully updated product", 201)

}
