package handlers

// we start with the login n verifing inputs and set the cookie, and we think about sorting files later

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User struct to handle login and registration data

type UserLogin struct {
	NameEmail string `json:"name_or_email"`
	Password  string `json:"password"`
}

// LoginHandler processes login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Authenticate user (check credentials from a database)
	/*if user.Username != "" && user.Email != "" && user.Password != "" {
		response := map[string]string{"message": "Login successful"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		}*/
	fmt.Println("login", user.NameEmail, "password", user.Password)
}
