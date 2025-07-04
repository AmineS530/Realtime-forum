package auth

// we start with the login n verifing inputs and set the cookie, and we think about sorting files later

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	helpers "RTF/back-end"
)

// User struct to handle login and registration data

type UserLogin struct {
	Name_Email string `json:"name_or_email"`
	Password   string `json:"password"`
}

// LoginHandler processes login requests
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.JsRespond(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var user UserLogin
	var userID int
	start := time.Now()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helpers.JsRespond(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if user.Name_Email != "" {
		userID, err = getID(user.Name_Email)
		if err != nil {
			helpers.JsRespond(w, err.Error(), http.StatusUnauthorized)
			return
		}
	}
	if CheckPassword(user.Password, userID) {
		jwtStart := time.Now()
		authorize(w, userID)
		helpers.InfoLog.Printf("jwt took %v", time.Since(jwtStart))
		helpers.InfoLog.Printf("Total login process took %v", time.Since(start))

		// todo: remove alerts to replace by js notifications
		helpers.JsRespond(w, "Login successful", http.StatusPermanentRedirect)
	} else {
		helpers.JsRespond(w, "Invalid Login info", http.StatusUnauthorized)
	}
}

func getID(nameOrEmail string) (int, error) {
	var isUsername bool
	lowCase := strings.ToLower(nameOrEmail)
	var userID int
	if _, exists := helpers.EntryExists("username", lowCase, "users", true); exists {
		isUsername = true
	} else if _, exists := helpers.EntryExists("email", lowCase, "users", true); exists {
		isUsername = false
	} else {
		return -1, fmt.Errorf("invalid Login info")
	}
	if isUsername {
		userID = int(getElemVal("id", "users", `LOWER(username) = "`+lowCase+`"`).(int64))
	} else {
		userID = int(getElemVal("id", "users", `LOWER(email) = "`+lowCase+`"`).(int64))
	}
	return userID, nil
}
