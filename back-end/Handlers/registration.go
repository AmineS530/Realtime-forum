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
	Age             int    `json:"age,string"`
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

	if !validInfo(w, user) {
		respondWithError(w, "Registration failed", http.StatusBadRequest)
		return
	}

	// use bcrypt to hash the password and store it in the database

	// store cerdentials and infos in different tables

	fmt.Printf("Registration data:\nusername: %s\nemail: %s\npassword: %s\npassword confirm: %s\nage: %d\ngender: %s\nfirst name: %s\nlast name: %s\n", user.Username, user.Email, user.Password, user.PasswordConfirm, user.Age, user.Gender, user.FirstName, user.LastName)
	// TODO: Save user to database
	response := map[string]string{"message": "Registration successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func validInfo(w http.ResponseWriter, user UserReg) bool {
	// check if username is valid and exists in the db
	if !isValidUsername(user.Username) {
		respondWithError(w, "Invalid username", http.StatusBadRequest)
		return false
	}
	if entryExists("username", user.Username, true) {
		respondWithError(w, "Username already exists", http.StatusBadRequest)
		return false
	}
	// check if password follows policy
	if !isValidPassword(user.Password) {
		respondWithError(w, "Invalid password format", http.StatusBadRequest)
		return false
	}

	// check if passwords match
	if user.Password != user.PasswordConfirm {
		respondWithError(w, "Passwords do not match", http.StatusBadRequest)
		return false
	}

	// // check if email is valid and exists in the db
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		respondWithError(w, "Invalid email format", http.StatusBadRequest)
		return false
	}
	if entryExists("email", user.Email, true) {
		respondWithError(w, "Email already exists", http.StatusBadRequest)
		return false
	}

	// check if age is between 15 and 90
	if user.Age < 15 || user.Age > 90 {
		respondWithError(w, "Invalid age", http.StatusBadRequest)
		return false
	}

	// check if gender is valid
	if user.Gender != "male" && user.Gender != "female" && user.Gender != "Attack Helicopter" {
		respondWithError(w, "Invalid option", http.StatusBadRequest)
		return false
	}

	// check if first name and last name are valid (no numbers and special characters)
	nameRegex := *regexp.MustCompile(`^[a-zA-Z ]+$`)
	if !nameRegex.MatchString(user.FirstName) || !nameRegex.MatchString(user.LastName) {
		respondWithError(w, "Invalid name format", http.StatusBadRequest)
		return false
	}

	return true
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

/*
func setToken() {
	token, err := uuid.NewV4()
	if err != nil {
		fmt.Println("Error generating token:", err)
	}

	query = `
		UPDATE users
		SET token = ?
		WHERE id = ?
	`
	_, err = db.Exec(query, token.String(), userInfo.ID)
	if err != nil {
		fmt.Println("Error storing token:", err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token.String(),
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	})
}
*/
