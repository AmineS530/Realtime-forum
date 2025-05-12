package requests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	helpers "RTF/back-end"
	"RTF/back-end/goFiles/auth"
)

type postInfo struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

func PostCreation(w http.ResponseWriter, r *http.Request, uid int) {
	if r.Method != http.MethodPost {
		auth.JsRespond(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var post postInfo
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		auth.JsRespond(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	categories := strings.Split(strings.Join(strings.Fields(post.Category), " "), ",")
	if len(categories) > 3 || len(categories) < 1 {
		auth.JsRespond(w, "Invalid category selection", http.StatusBadRequest)
		return
	}

	// Validate basic input
	if len(post.Title) < 3 || len(post.Content) < 10 {
		auth.JsRespond(w, "Title and content required", http.StatusBadRequest)
		return
	}
	if !postPost(post, categories, uid) {
		auth.JsRespond(w, "Post creation failed", http.StatusBadRequest)
	}
	auth.JsRespond(w, "Post created successfully", http.StatusOK)
}

func postPost(post postInfo, categories []string, uid int) bool {
	fmt.Println(post, categories, uid)
	query := `
	INSERT
		INTO posts
		(title, content, uid, categories)
	VALUES (?, ?, ?, ?) `
	_, err := helpers.DataBase.Exec(query,
		post.Title,
		post.Content,
		uid,
		strings.Join(categories, ", "))
	if err != nil {
		helpers.ErrorLog.Println("Database insertion error:", err)
		return false
	}
	return true
}
