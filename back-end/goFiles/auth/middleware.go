package auth

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"
)

type contextKey string

const UserContextKey contextKey = "user"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handleAuthError := func(w http.ResponseWriter, r *http.Request, message string) {
			if strings.Contains(r.Header.Get("Accept"), "text/html") {
				helpers.HtmlTemplates.ExecuteTemplate(w, "index.html", nil)
			} else {
				helpers.JsRespond(w, message, http.StatusUnauthorized)
			}
		}

		// Extract and verify JWT
		token, err := ExtractJWT(r)
		if err != nil {
			handleAuthError(w, r, "Missing authentication token")
			return
		}

		payload, err := jwt.JWTVerify(token)
		if err != nil {
			handleAuthError(w, r, "Invalid or expired token")
			return
		}
		// Verify session
		sessionID, err := ExtractSSID(r)
		if err != nil {
			handleAuthError(w, r, "Missing session")
			return
		}

		if !isValidSession(payload.Sub, sessionID) {
			handleAuthError(w, r, "Invalid session")
			return
		}

		// Set user in context and proceed
		ctx := context.WithValue(r.Context(), UserContextKey, payload)
		w.Header().Set("Content-Type", "application/json")
		next(w, r.WithContext(ctx))
	}
}

func ApiOnlyAccess(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAPIRequest := r.Header.Get("X-Requested-With") == "XMLHttpRequest" ||
			strings.Contains(r.Header.Get("Accept"), "application/json")

		if !isAPIRequest {
			if strings.Contains(r.Header.Get("Accept"), "text/html") {
				helpers.HtmlTemplates.ExecuteTemplate(w, "index.html", nil)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
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
