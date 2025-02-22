package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	helpers "RTF/back-end"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 16 {
		return false
	}

	hasDigit := regexp.MustCompile(`[0-9]`)
	hasLower := regexp.MustCompile(`[a-z]`)
	hasUpper := regexp.MustCompile(`[A-Z]`)
	hasSpecial := regexp.MustCompile(`[-#$.%&*]`)

	return hasDigit.MatchString(password) &&
		hasLower.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasSpecial.MatchString(password)
}

func isValidUsername(username string) bool {
	if len(username) < 5 || len(username) > 18 {
		return false
	}

	hasAlpha := regexp.MustCompile(`[a-zA-Z]`)
	hasNoSpaces := regexp.MustCompile(`^\S*$`)
	hasValidChars := regexp.MustCompile(`^[a-zA-Z0-9_-]*$`)

	return hasAlpha.MatchString(username) &&
		hasNoSpaces.MatchString(username) &&
		hasValidChars.MatchString(username)
}

func entryExists(elem, value string, checkLower bool) bool {
	var count int

	query := fmt.Sprintf("SELECT COUNT(*) FROM users WHERE %s = ?", elem)
	if checkLower {
		query = fmt.Sprintf("SELECT COUNT(*) FROM users WHERE LOWER(%s) = LOWER(?)", elem)
	}

	err := helpers.DataBase.QueryRow(query, value).Scan(&count)
	if err != nil {
		helpers.ErrorLog.Fatalln("Database error:", err)
		return false
	}

	return count > 0
}

func respondWithError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
