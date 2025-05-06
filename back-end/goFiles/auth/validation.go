package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"

	jwt "RTF/back-end/goFiles/JWT"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 30 {
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

func JsRespond(w http.ResponseWriter, message string, code int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		json.NewEncoder(w).Encode(ErrorResponse{Error: message})
	}
}

func VerifyUser(jwt_token, session_id string) (bool, error) {
	if session_id == "" {
		return false, errors.New("missing session_id")
	}

	payload, err := jwt.JWTVerify(jwt_token)
	if err != nil || payload == nil {
		return false, errors.New("invalid or expired token")
	}

	valid, _ := SessionExists(payload.Sub, session_id)
	if !valid {
		return false, errors.New("invalid session")
	}

	return true, nil
}
