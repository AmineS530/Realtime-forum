package lab

import (
	"fmt"
	"net/http"

	jwt "RTF/back-end/goFiles/JWT"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Get user data from context
	if user, ok := r.Context().Value("user").(*jwt.JwtPayload); ok {
		// Step 2: Use user data (e.g., print username)
		fmt.Println("Logged in user:", user.Username)
		// You can now access other user data like user.ID, user.Email, etc.
	} else {
		http.Error(w, "User not found in context", http.StatusUnauthorized)
	}
}
