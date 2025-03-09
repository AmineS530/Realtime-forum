package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	helpers "RTF/back-end"

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
	// store cerdentials and infos in different tables
	if !insertInfo(user) {
		respondWithError(w, "Registration failed", http.StatusInternalServerError)
		return
	}
	// get the id of the inserted values

	// use bcrypt to hash the password and store it in the database
	userID := int(getElemVal("id", "users", `username = "`+user.Username+`"`).(int64))
	changePassword(user.Password, userID)

	authorize(w, userID)
	respondWithError(w, "Login successful", http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func validInfo(w http.ResponseWriter, user UserReg) bool {
	// check if username is valid and exists in the db
	if !isValidUsername(user.Username) {
		respondWithError(w, "Invalid username", http.StatusBadRequest)
		return false
	}
	if _, exists := helpers.EntryExists("username", user.Username, "users", true); exists {
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

	// check if email is valid and exists in the db
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(user.Email) {
		respondWithError(w, "Invalid email format", http.StatusBadRequest)
		return false
	}
	if _, exists := helpers.EntryExists("email", user.Email, "users", true); exists {
		respondWithError(w, "Email already exists", http.StatusBadRequest)
		return false
	}

	// check if age is between 15 and 90
	if user.Age < 15 || user.Age > 90 {
		respondWithError(w, "Invalid age", http.StatusBadRequest)
		return false
	}

	// check if gender is valid
	if user.Gender != "male" && user.Gender != "female" && user.Gender != "Attack helicopter" {
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

func insertInfo(user UserReg) bool {
	query := `INSERT INTO users (username, email, age, gender, first_name, last_name) 
	          VALUES (?, ?, ?, ?, ?, ?)`

	_, err := helpers.DataBase.Exec(query,
		user.Username,
		user.Email,
		user.Age,
		user.Gender,
		user.FirstName,
		user.LastName,
	)
	if err != nil {
		helpers.ErrorLog.Fatalln("Database insertion error:", err)
		return false
	}
	return true
}

func changePassword(password string, userID int) {
	query := `INSERT INTO credentials (id, hash) VALUES (?, ?)`
	_, err := helpers.DataBase.Exec(query, userID, HashPassword(password))
	if err != nil {
		helpers.ErrorLog.Fatalln(err.Error())
	}
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		helpers.ErrorLog.Fatalln(err.Error())
		return ""
	}
	return string(bytes)
}

func CheckPassword(password string, userID int) bool {
	var hashedPassword string

	query := `SELECT hash FROM credentials WHERE id = ?`
	err := helpers.DataBase.QueryRow(query, userID).Scan(&hashedPassword)
	if err != nil {
		helpers.ErrorLog.Fatalln(err.Error())
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

func getElemVal(selectedElem, from, where string) any {
	var res any

	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s", selectedElem, from, where)

	err := helpers.DataBase.QueryRow(query).Scan(&res)
	if err != nil {
		if err == sql.ErrNoRows {
			res = ""
		} else {
			helpers.ErrorLog.Println("Database error:", err)
		}
	}

	return res
}
