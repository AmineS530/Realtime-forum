package auth

import (
	"errors"
	"regexp"
	"unicode"

	jwt "RTF/back-end/goFiles/JWT"
)

func hasSpecial(s string) bool {
	for _, ch := range s {
		if unicode.IsPunct(ch) || unicode.IsSymbol(ch) {
			return true
		}
	}
	return false
}

func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 30 {
		return false
	}

	hasDigit := regexp.MustCompile(`[0-9]`)
	hasLower := regexp.MustCompile(`[a-z]`)
	hasUpper := regexp.MustCompile(`[A-Z]`)

	return hasDigit.MatchString(password) &&
		hasLower.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasSpecial(password)
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
