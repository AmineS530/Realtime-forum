package handlers

// we start with the login n verifing inputs and set the cookie, and we think about sorting files later

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// User struct to handle login and registration data

type UserLogin struct {
	Name_Email string `json:"name_or_email"`
	Password   string `json:"password"`
}

// LoginHandler processes login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user UserLogin
	var userID int
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Authenticate user (check credentials from a database)
	if user.Name_Email != "" {
		userID, err = getID(user.Name_Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
	} 
	if CheckPassword(user.Password, userID) {
		authorize(userID, w)
		response := map[string]string{"message": "Login successful"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	} else {
		http.Error(w, "Invalid Login info", http.StatusUnauthorized)
	}
	fmt.Printf("Login data:\nuser: %s\npassword: %s\n", user.Name_Email, user.Password)
}

func getID(nameOrEmail string) (int, error) {
	var isUsername bool
	var userID int
	if entryExists("username", nameOrEmail, true) {
		isUsername = true
	} else if entryExists("email", nameOrEmail, true) {
		isUsername = false
	} else {
		return -1, fmt.Errorf("Invalid Login info")
	}
	if isUsername {
		userID = int(getElemVal("id", "users", `username = "`+nameOrEmail+`"`).(int64))
	} else {
		userID = int(getElemVal("id", "users", `email = "`+nameOrEmail+`"`).(int64))
	}
	return userID, nil
}
