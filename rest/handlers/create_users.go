package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	var newUser database.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	createdUser := newUser.Store()
	utils.SendData(w, createdUser, http.StatusCreated)

}
