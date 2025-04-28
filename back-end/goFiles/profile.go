package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	jwt "RTF/back-end/goFiles/JWT"
	"RTF/back-end/goFiles/auth"
	"RTF/global"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	Age       int    `json:"age"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Email     string `json:"email"`
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, `{"error": "Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}
	tok, _ := auth.ExtractJWT(r)
	payload, _ := jwt.JWTVerify(tok)
	username := payload.Username

	userInfo := user{}
	query := `SELECT age, username, first_name, last_name, gender, email FROM users WHERE username = ?`

	err := global.DataBase.QueryRow(query, username).Scan(&userInfo.Age, &userInfo.Username, &userInfo.FirstName, &userInfo.LastName, &userInfo.Gender, &userInfo.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, `{"error": "User not found"}`, http.StatusNotFound)
		} else {
			http.Error(w, `{"error": "`+err.Error()+`"}`, http.StatusInternalServerError)
		}
		return
	}

	// Convert struct to map
	profile := map[string]interface{}{
		"age":        userInfo.Age,
		"username":   userInfo.Username,
		"first_name": userInfo.FirstName,
		"last_name":  userInfo.LastName,
		"gender":     userInfo.Gender,
		"email":      userInfo.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	// get user info from request body
	// validate user info
	// update user info in database
	// respond with success or error message
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	// get user id from request
	// delete user from database
	// respond with success or error message
}
