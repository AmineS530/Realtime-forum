package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/dms"
	"RTF/back-end/goFiles/requests"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if helpers.Err != nil {
		auth.JsRespond(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		auth.JsRespond(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/api/") {
		auth.JsRespond(w, "API endpoint not found", http.StatusNotFound)
		return
	}
	if r.URL.Path == "/" || r.URL.Path == "/profile" {
		if err := helpers.HtmlTemplates.ExecuteTemplate(w, "index.html", nil); err != nil {
			fmt.Println("Error executing template: ", err.Error())
			auth.JsRespond(w, "Error executing template", http.StatusInternalServerError)
			return
		}
	} else {
		// 	// helpers.ErrorPagehandler(w, http.StatusNotFound)
		auth.JsRespond(w, "Page not found", http.StatusNotFound)
		return
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r.URL)
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
		payload := r.Context().Value(auth.UserContextKey)
		data, ok := payload.(*jwt.JwtPayload)
		if !ok {
			helpers.ErrorPagehandler(w, http.StatusBadRequest)
			fmt.Println("azer qsdf")
			return
		}
		usernames, _ := dms.GetUserNames(data.Sub)
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

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		auth.JsRespond(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Println(r.URL)
	jwtToken, _ := auth.ExtractJWT(r)
	payload, _ := jwt.JWTVerify(jwtToken)
	x := r.PathValue("type")
	fmt.Println("here", x)
	switch x {
	case "createPost":
		requests.PostCreation(w, r, payload.Sub)

	case "createComment":
		requests.CommentCreation(w, r, payload.Sub)
	}
}
