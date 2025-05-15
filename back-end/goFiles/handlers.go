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
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if helpers.Err != nil {
		helpers.JsRespond(w, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		helpers.JsRespond(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := helpers.HtmlTemplates.ExecuteTemplate(w, "index.html", nil); err != nil {
		fmt.Println("Error executing template: ", err.Error())
		helpers.JsRespond(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
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
			helpers.JsRespond(w, "Invalid payload", http.StatusBadRequest)
			return
		}
		usernames, _ := dms.GetUserNames(data.Sub)
		jsonData, _ := json.Marshal(usernames)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	case "dmhistory":
		target := r.Header.Get("target")
		page := r.Header.Get("page")
		tok, _ := auth.ExtractJWT(r)
		payload, _ := jwt.JWTVerify(tok)
		username := payload.Username
		dms, err := dms.GetdmHistory(username, target, page)
		if err != nil {
			helpers.ErrorLog.Print("routes.go 69", err)
		}
		jsonData, _ := json.Marshal(dms)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	default:
		helpers.JsRespond(w, "Invalid request type", http.StatusBadRequest)
		return
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		helpers.JsRespond(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	jwtToken, _ := auth.ExtractJWT(r)
	payload, _ := jwt.JWTVerify(jwtToken)
	x := r.PathValue("type")
	switch x {
	case "createPost":
		requests.PostCreation(w, r, payload.Sub)
	case "createComment":
		requests.CommentCreation(w, r, payload.Sub)
	}
}
