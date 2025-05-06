package handlers

import (
	"net/http"
	"strings"

	helpers "RTF/back-end"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/ws"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/api/v1/ws", ws.HandleConnections)
	mux.HandleFunc("/api/v1/get/{type}", auth.AuthMiddleware(auth.ApiOnlyAccess(GetHandler)))
	mux.HandleFunc("/api/v1/post/{type}", auth.AuthMiddleware(auth.ApiOnlyAccess(PostHandler)))
	mux.HandleFunc("/api/profile", auth.AuthMiddleware(auth.ApiOnlyAccess(ProfileHandler)))
	ProtectedStatic(mux, "/front-end/styles/", "./front-end/styles/")
	ProtectedStatic(mux, "/front-end/scripts/", "./front-end/scripts/")
	ProtectedStatic(mux, "/front-end/images/", "./front-end/images/")
	mux.HandleFunc("/profile", auth.AuthMiddleware(IndexHandler))
	mux.HandleFunc("/api/check-auth", auth.ApiOnlyAccess(auth.CheckAuthHandler))
	mux.HandleFunc("/api/login", auth.LoginHandler)
	mux.HandleFunc("/api/register", auth.RegisterHandler)
	mux.HandleFunc("/api/logout", auth.Logout)
	helpers.ServerRoutine()

	return mux
}

func ProtectedStatic(mux *http.ServeMux, routePrefix, dirPath string) {
	fs := http.FileServer(http.Dir(dirPath))

	mux.HandleFunc(routePrefix, func(w http.ResponseWriter, r *http.Request) {
		ref := r.Referer()
		if ref == "" || !strings.Contains(ref, r.Host) {
			http.Error(w, "Direct access forbidden", http.StatusForbidden)
			return
		}
		// Strip the prefix correctly without trimming the trailing slash first
		http.StripPrefix(routePrefix, fs).ServeHTTP(w, r)
	})
}
