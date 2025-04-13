package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	helpers "RTF/back-end"
)

func authorize(w http.ResponseWriter, userID int) {
	username := getElemVal("username", "users", `id = "`+strconv.Itoa(userID)+`"`).(string)
	jwt, sessionID, err := CheckSession(userID, username)
	if err != nil {
		helpers.ErrorLog.Println(err)
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt",
		Value:    jwt,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(60 * time.Minute),
	})

	// Set the Session ID in a separate HttpOnly cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "ssid",
		Value:    sessionID,
		Path:     "/",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		Expires:  time.Now().Add(60 * time.Minute),
	})

	w.WriteHeader(http.StatusOK)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// todo check the http.errors
	sessionCookie, _ := ExtractSSID(r)
	invalidateSession(sessionCookie)
	http.SetCookie(w, &http.Cookie{
		Name:    "jwt",
		Value:   "",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "ssid",
		Value:   "",
		Path:    "/",
		MaxAge:  -1,
		Expires: time.Unix(0, 0),
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CheckAuthHandler(w http.ResponseWriter, r *http.Request) {
	// check online from js based on the api
	jwt_token, _ := ExtractJWT(r)
	ssid, _ := ExtractSSID(r)
	auth, err := VerifyUser(jwt_token, ssid)
	count, _ := helpers.EntryExists("session_id", ssid, "sessions", false)
	if count != 1 {
		fmt.Println("Session ID not found in database")
		http.SetCookie(w, &http.Cookie{
			Name:    "jwt",
			Value:   "",
			Path:    "/",
			MaxAge:  -1,
			Expires: time.Unix(0, 0),
		})
		http.SetCookie(w, &http.Cookie{
			Name:    "ssid",
			Value:   "",
			Path:    "/",
			MaxAge:  -1,
			Expires: time.Unix(0, 0),
		})
	}

	if !auth || err != nil {
		fmt.Println("Error verifying user:", err)
		json.NewEncoder(w).Encode(map[string]bool{"authenticated": false})
		return
	}

	json.NewEncoder(w).Encode(map[string]bool{"authenticated": true})
}
