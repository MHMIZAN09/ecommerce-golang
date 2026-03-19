package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {

	utils.SendData(w, database.Products, 200)
}
