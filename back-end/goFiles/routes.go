package handlers

import (
	"fmt"
	"net/http"
	"strings"

	helpers "RTF/back-end"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/ws"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/api/v1/get/{type}", auth.AuthMiddleware(dumbjson))
	mux.HandleFunc("/api/profile", auth.AuthMiddleware(ProfileHandler))
	mux.HandleFunc("/api/ws", ws.HandleWebSocket)
	mux.Handle("/front-end/styles/", http.StripPrefix("/front-end/styles/", http.FileServer(http.Dir("./front-end/styles"))))
	mux.Handle("/front-end/scripts/", http.StripPrefix("/front-end/scripts/", http.FileServer(http.Dir("./front-end/scripts"))))
	mux.Handle("/front-end/images/", http.StripPrefix("/front-end/images/", http.FileServer(http.Dir("./front-end/images"))))
	mux.HandleFunc("/profile", auth.AuthMiddleware(IndexHandler))

	mux.HandleFunc("/api/check-auth", auth.CheckAuthHandler)
	mux.HandleFunc("/api/login", auth.LoginHandler)
	mux.HandleFunc("/api/register", auth.RegisterHandler)
	mux.HandleFunc("/api/logout", auth.Logout)
	helpers.ServerRoutine()

	return mux
}

// TODO sMArT
func dumbjson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	x := r.PathValue("type")

	// ErrorPagehandler(w, http.StatusInternalServerError, "azerqsdfwxcv")
	if x == "comments" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(commentjson))
	} else if x == "posts" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(postjson))
	} else {
		helpers.ErrorPagehandler(w, http.StatusNotFound)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if helpers.Err != nil {
		helpers.ErrorPagehandler(w, http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/api/") {
		helpers.ErrorPagehandler(w, http.StatusNotFound)
		return
	}
	if r.URL.Path == "/" || r.URL.Path == "/profile" {
		if err := helpers.HtmlTemplates.ExecuteTemplate(w, "index.html", nil); err != nil {
			fmt.Println("Error executing template: ", err.Error())
			helpers.ErrorPagehandler(w, http.StatusInternalServerError)
			return
		}
	} else {
		helpers.ErrorPagehandler(w, http.StatusNotFound)
		return
	}
}
