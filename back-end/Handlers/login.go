package handlers

// we start with the login n verifing inputs and set the cookie, and we think about sorting files later

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User struct to handle login and registration data
type UserReg struct {
	Username        string `json:"username,omitempty"`
	Email           string `json:"email,omitempty"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation,omitempty"`
	Age             int    `json:"age,omitempty"`
	Gender          string `json:"gender,omitempty"`
	FirstName       string `json:"first_name,omitempty"`
	LastName        string `json:"last_name,omitempty"`
}

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

// RegisterHandler processes user registrations
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user UserReg
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Basic validation
	if user.Password != user.PasswordConfirm {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// TODO: Save user to database
	response := map[string]string{"message": "Registration successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
