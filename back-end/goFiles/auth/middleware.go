package auth

import (
	"context"
	"errors"
	"log"
	"net/http"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"
	"RTF/global"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := ExtractJWT(r)
		if err != nil {
			helpers.ErrorPagehandler(w, http.StatusUnauthorized)
			// json.NewEncoder(w).Encode(map[string]string{"error": "Missing JWT"})
			return
		}

		// Verify JWT token
		payload, err := jwt.JWTVerify(token)
		if err != nil {
			helpers.ErrorPagehandler(w, http.StatusUnauthorized)
			// json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JWT"})
			return
		}
		sessionID, err := ExtractSSID(r)
		if err != nil {
			helpers.ErrorPagehandler(w, http.StatusUnauthorized)
			// json.NewEncoder(w).Encode(map[string]string{"error": "Missing session"})
			return
		}

		// Validate session ID from database
		isValidSession := isValidSession(payload.Sub, sessionID)
		if !isValidSession {
			helpers.ErrorPagehandler(w, http.StatusUnauthorized)
			// json.NewEncoder(w).Encode(map[string]string{"error": "Invalid session"})
			return
		}
		ctx := context.WithValue(r.Context(), UserContextKey, payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Extract JWT from Authorization header or cookie
func ExtractJWT(r *http.Request) (string, error) {
	cookie, err := r.Cookie("jwt")
	if err != nil {
		return "", errors.New("jwt cookie not found")
	}

	return cookie.Value, nil
}

// Extract session ID from cookie
func ExtractSSID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("ssid")
	if err != nil {
		return "", errors.New("session_id cookie not found")
	}
	return cookie.Value, nil
}

// Validate session in database
func isValidSession(userID int, sessionID string) bool {
	var count int
	err := global.DataBase.QueryRow("SELECT COUNT(*) FROM sessions WHERE user_id = ?", userID).Scan(&count)
	if count > 1 || err != nil {
		return false
	}
	err = global.DataBase.QueryRow("SELECT COUNT(*) FROM sessions WHERE user_id = ? AND session_id = ?", userID, sessionID).Scan(&count)
	if err != nil {
		log.Println("Error checking session:", err)
		return false
	}
	return count == 1
}
