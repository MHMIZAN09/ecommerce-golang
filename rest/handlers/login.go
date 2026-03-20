package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login (w http.ResponseWriter, r *http.Request) {
	var reqLogin ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user := database.FindUserByEmailAndPassword(reqLogin.Email, reqLogin.Password)
	if user == nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}
	

	utils.SendData(w, user, http.StatusOK)
}
