package handlers

import (
	"net/http"

	helpers "RTF/back-end"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/ws"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/api/v1/ws", ws.HandleConnections)
	mux.HandleFunc("/api/v1/get/{type}", auth.AuthMiddleware(GetHandler))
	mux.HandleFunc("/api/v1/post/{type}", auth.AuthMiddleware(PostHandler))
	mux.HandleFunc("/api/profile", auth.AuthMiddleware(ProfileHandler))
	ProtectedStatic(mux, "/front-end/styles/", "./front-end/styles")
	ProtectedStatic(mux, "/front-end/scripts/", "./front-end/scripts")
	ProtectedStatic(mux, "/front-end/images/", "./front-end/images")
	mux.HandleFunc("/profile", auth.AuthMiddleware(IndexHandler))
	mux.HandleFunc("/api/check-auth", auth.CheckAuthHandler)
	mux.HandleFunc("/api/login", auth.LoginHandler)
	mux.HandleFunc("/api/register", auth.RegisterHandler)
	mux.HandleFunc("/api/logout", auth.Logout)
	helpers.ServerRoutine()

	return mux
}

func ProtectedStatic(mux *http.ServeMux, routePrefix string, dirPath string) {
	mux.HandleFunc(routePrefix, func(w http.ResponseWriter, r *http.Request) {
		if !BlockDirectAccess(w, r) {
			return
		}
		http.StripPrefix(routePrefix, http.FileServer(http.Dir(dirPath))).ServeHTTP(w, r)
	})
}

func BlockDirectAccess(w http.ResponseWriter, r *http.Request) bool {
	if r.Referer() == "" {
		helpers.ErrorPagehandler(w, http.StatusForbidden)
		return false
	}
	return true
}
