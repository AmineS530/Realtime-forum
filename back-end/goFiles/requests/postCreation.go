package requests

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var post postInfo
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	categories := strings.Split(post.Category, " ")
	if len(categories) > 3 || len(categories) < 1 {
		http.Error(w, "Invalid category selection", http.StatusBadRequest)
		return
	}

	// Validate basic input
	if len(post.Title) < 3 || len(post.Content) < 10 {
		http.Error(w, "Title and content required", http.StatusBadRequest)
		return
	}
	postPost(post, categories, uid)
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

func CommentCreation(w http.ResponseWriter, r *http.Request) {
	
}