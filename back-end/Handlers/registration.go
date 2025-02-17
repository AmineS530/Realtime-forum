package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
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

	/*
		password policy and regex ^(?=.*\d)(?=.*[a-zA-Z])(?=.*[A-Z])(?=.*[-\#\$\.\%\&\*])(?=.*[a-zA-Z]).{8,16}$
			At least 8 - 16 characters,
			must contain at least 1 uppercase letter,
			must contain at least 1 lowercase letter,
			and 1 number
			Can contain any of this special characters $ % # * & - .
	*/
	// check if poassword follow policy
	if isValidPassword(user.Password) {
		http.Error(w, "Invalid password format", http.StatusBadRequest)
		return
	}
	// check if passwords match
	if user.Password != user.PasswordConfirm {
		http.Error(w, "Passwords do not match", http.StatusBadRequest)
		return
	}

	// check if username is valid and exists in the db

	// check if email is valid and exists in the db
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		http.Error(w, "Invalid email format", http.StatusBadRequest)
	}

	// check if age is between 15 and 90
	if user.Age < 15 || user.Age > 90 {
		http.Error(w, "Invalid age", http.StatusBadRequest)
	}
	// check if gender is valid
	if user.Gender != "male" && user.Gender != "female" && user.Gender != "Attack Helicopter" {
		http.Error(w, "Invalid gender", http.StatusBadRequest)
	}
	// check if first name and last name are valid (no numbers and special characters)
	nameRegex := *regexp.MustCompile(`^[a-z A-Z]+$`)
	if !nameRegex.MatchString(user.FirstName) || !nameRegex.MatchString(user.LastName) {
		http.Error(w, "Invalid name format", http.StatusBadRequest)
	}
	// use bcrypt to hash the password and store it in the database

	// store cerdentials and infos in different tables

	fmt.Printf("Registration data:\nusername: %s\nemail: %s\npassword: %s\npassword confirm: %s\nage: %d\ngender: %s\nfirst name: %s\nlast name: %s\n", user.Username, user.Email, user.Password, user.PasswordConfirm, user.Age, user.Gender, user.FirstName, user.LastName)
	// TODO: Save user to database
	response := map[string]string{"message": "Registration successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func isValidPassword(password string) bool {
	// Ensure length is between 8 and 16 characters
	if len(password) < 8 || len(password) > 16 {
		return false
	}

	// Define separate regexes for each condition
	hasDigit := regexp.MustCompile(`[0-9]`)
	hasLower := regexp.MustCompile(`[a-z]`)
	hasUpper := regexp.MustCompile(`[A-Z]`)
	hasSpecial := regexp.MustCompile(`[-#$.%&*]`)

	// Check all conditions
	return hasDigit.MatchString(password) &&
		hasLower.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasSpecial.MatchString(password)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 16)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func CheckPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
