package auth

import (
	"errors"
	"net/http"
)

// still need to put the middleware logic to retun the user to login page, idk how it works with the singlepage app

var Middleware = []func(http.HandlerFunc) http.HandlerFunc{
	authMiddleware,
	loginMiddleware,
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your logic here
		next.ServeHTTP(w, r)
	})
}

func loginMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your logic here
		next.ServeHTTP(w, r)
	})
}

// Context key for user info
// type contextKey string

// const UserContextKey contextKey = "user"

// // AuthMiddleware checks JWT and Session ID
// func AuthMiddleware(db *sql.DB) func(http.Handler) http.Handler {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			// Extract JWT from Authorization header or cookie
// 			token := extractJWT(r)
// 			if token == "" {
// 				http.Error(w, "Unauthorized: Missing token", http.StatusUnauthorized)
// 				return
// 			}

// 			// Verify JWT
// 			payload, err := jwt.JWTVerify(token)
// 			if err != nil {
// 				http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
// 				return
// 			}

// 			// Check session from cookie
// 			sessionID, err := extractSessionID(r)
// 			if err != nil {
// 				http.Error(w, "Unauthorized: Missing session", http.StatusUnauthorized)
// 				return
// 			}

// 			// Validate session in DB
// 			if !isValidSession(db, , sessionID) {
// 				http.Error(w, "Unauthorized: Invalid session", http.StatusUnauthorized)
// 				return
// 			}

// 			// Attach user info to request context
// 			ctx := context.WithValue(r.Context(), UserContextKey, payload)
// 			next.ServeHTTP(w, r.WithContext(ctx))
// 		})
// 	}
// }

// Extract JWT from Authorization header or cookie
func ExtractJWT(r *http.Request) string {
	cookie, err := r.Cookie("jwt")
	if err == nil {
		return cookie.Value
	}

	return ""
}

// Extract session ID from cookie
func extractSessionID(r *http.Request) (string, error) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", errors.New("session_id cookie not found")
	}
	return cookie.Value, nil
}

// // Validate session in database
// func isValidSession(db *sql.DB, userID int, sessionID string) bool {
// 	var count int
// 	err := db.QueryRow("SELECT COUNT(*) FROM sessions WHERE user_id = ? AND session_id = ?", userID, sessionID).Scan(&count)
// 	if err != nil {
// 		log.Println("Error checking session:", err)
// 		return false
// 	}
// 	return count == 1
// }
