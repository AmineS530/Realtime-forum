package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/dms"
	"RTF/back-end/goFiles/requests"
	"RTF/back-end/goFiles/ws"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/api/v1/ws", ws.HandleConnections)
	mux.HandleFunc("/api/v1/get/{type}", auth.AuthMiddleware(dumbjson))
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

// TODO sMArT
func dumbjson(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	x := r.PathValue("type")
	offset := r.URL.Query().Get("offset")
	switch x {
	case "comments":
		pid := r.URL.Query().Get("pid")
		comments, _ := requests.GetComments(pid)
		jsoncomment, _ := json.Marshal(comments)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(jsoncomment))
	case "posts":
		posts, _ := requests.GetPosts(offset)
		jsonData, _ := json.Marshal(posts)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	case "users":
		usernames, _ := dms.GetUserNames()
		jsonData, _ := json.Marshal(usernames)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	case "dmhistory":
		target := r.Header.Get("target")
		tok, _ := auth.ExtractJWT(r)
		payload, _ := jwt.JWTVerify(tok)
		username := payload.Username
		dms, err := dms.GetdmHistory(username, target)
		if err != nil {
			helpers.ErrorLog.Print("routes.go 69", err)
		}
		jsonData, _ := json.Marshal(dms)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
		fmt.Println(target, username, dms)
	default:
		helpers.ErrorPagehandler(w, http.StatusNotFound)
		return
	}
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
