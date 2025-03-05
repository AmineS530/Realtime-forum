package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"
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

func respondWithError(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}

func VerifyUser(jwtToken, session_id string) (bool, error) {
	payload, err := jwt.JWTVerify(jwtToken)
	if err != nil {
		helpers.ErrorLog.Println(err)
		return false, err
	}
	if session_id != "" {
		if count, _ := helpers.EntryExists("session_id", session_id, "sessions", true); count != 1 {
			InvalidateSession(int(payload.Sub))
			return false, errors.New("invalid session")
		}
	}

	if count, _ := helpers.EntryExists("username", payload.Username, "users", true); count != 1 {
		return false, errors.New("invalid username")
	}
	return true, nil
}
