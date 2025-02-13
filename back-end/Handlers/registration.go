package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// verifying the input from the json api and the registeration logic
type UserReg struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
	Age             int    `json:"age,string"` // <-- Fix: Accept string and convert to int
	Gender          string `json:"gender"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
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
	// if user.Password != user.PasswordConfirm {
	// 	http.Error(w, "Passwords do not match", http.StatusBadRequest)
	// 	return
	// }
	fmt.Println("user", user.Username, "email", user.Email, "password", user.Password, "password_confirm", user.PasswordConfirm, "age", user.Age, "gender", user.Gender, "first_name", user.FirstName, "last_name", user.LastName)

	// TODO: Save user to database
	response := map[string]string{"message": "Registration successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
