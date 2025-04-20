package auth

import (
	"encoding/json"
	"errors"
	"net/http"
	"regexp"
	"strconv"

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

func JsRespond(w http.ResponseWriter, message string, code int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ErrorResponse{Error: message})
	}
	// w.WriteHeader(code) why not ??     NOTE : when registration erors respond with 200 status
}

func VerifyUser(jwt_token, session_id string) (bool, error) {
	userSessionCount := 0
	payload, _ := jwt.JWTVerify(jwt_token)
	if payload == nil || session_id == "" {
		return false, errors.New("missing payload or session_id")
	}
	if session_id != "" {
		if userSessionCount, _ = helpers.EntryExists("user_id", strconv.Itoa(int(payload.Sub)), "sessions", true); userSessionCount > 1 {
			err := InvalidateSessions(payload.Sub)
			if err != nil {
				return false, err
			}
			return false, errors.New("invalid session_id")
		}
	}
	if count, _ := helpers.EntryExists("username", payload.Username, "users", true); count != 1 {
		return false, errors.New("invalid username")
	}
	return true, nil
}
