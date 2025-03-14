package auth

import (
	"context"
	"errors"
	"log"
	"net/http"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"
)

type contextKey string

const UserContextKey contextKey = "user"

// todo send req to js to c lear cookies
var Middleware = []func(http.HandlerFunc) http.HandlerFunc{
	// authMiddleware,
	// loginMiddleware,
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := ExtractJWT(r)
		if err != nil {
			http.Error(w, "Unauthorized: Missing JWT", http.StatusUnauthorized)
			return
		}

		// Verify JWT token
		payload, err := jwt.JWTVerify(token)
		if err != nil {
			http.Error(w, "Unauthorized: Invalid JWT", http.StatusUnauthorized)
			return
		}
		sessionID, err := ExtractSSID(r)
		if err != nil {
			http.Error(w, "Unauthorized: Missing session", http.StatusUnauthorized)
			return
		}

		// Validate session ID from database
		isValidSession := isValidSession(payload.Sub, sessionID)
		if !isValidSession {
			http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), UserContextKey, payload)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := ExtractJWT(r)
		if err != nil {
			http.Error(w, "Unauthorized: Not logged in", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
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
	err := helpers.DataBase.QueryRow("SELECT COUNT(*) FROM sessions WHERE user_id = ?", userID).Scan(&count)
	if count > 1 || err != nil {
		return false
	}
	err = helpers.DataBase.QueryRow("SELECT COUNT(*) FROM sessions WHERE user_id = ? AND session_id = ?", userID, sessionID).Scan(&count)
	if err != nil {
		log.Println("Error checking session:", err)
		return false
	}
	return count == 1
}
