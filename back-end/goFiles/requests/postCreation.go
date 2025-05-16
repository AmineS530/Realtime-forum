package requests

import (
	"encoding/json"
	"html"
	"net/http"
	"strings"

	helpers "RTF/back-end"
)

type postInfo struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

func PostCreation(w http.ResponseWriter, r *http.Request, uid int) {
	if r.Method != http.MethodPost {
		helpers.JsRespond(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var post postInfo
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		helpers.JsRespond(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	categories := strings.Split(strings.Join(strings.Fields(html.EscapeString(post.Category)), " "), ",")
	if len(categories) > 3 || len(categories) < 1 {
		helpers.JsRespond(w, "Invalid category selection", http.StatusBadRequest)
		return
	}

	// Validate basic input
	if len(post.Title) < 3 || len(post.Content) < 10 {
		helpers.JsRespond(w, "Title and content required", http.StatusBadRequest)
		return
	}
	if !postPost(post, categories, uid) {
		helpers.JsRespond(w, "Post creation failed", http.StatusBadRequest)
	}
	helpers.JsRespond(w, "Post created successfully", http.StatusOK)
}

func postPost(post postInfo, categories []string, uid int) bool {
	query := `
	INSERT
		INTO posts
		(title, content, uid, categories)
	VALUES (?, ?, ?, ?) `
	_, err := helpers.DataBase.Exec(query,
		html.EscapeString(post.Title),
		html.EscapeString(post.Content),
		uid,
		strings.Join(categories, ", "))
	if err != nil {
		helpers.ErrorLog.Println("Database insertion error:", err)
		return false
	}
	return true
}
