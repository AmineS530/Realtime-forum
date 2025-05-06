package requests

import (
	"encoding/json"
	"net/http"

	helpers "RTF/back-end"
	"RTF/back-end/goFiles/auth"
)

type commentInfo struct {
	PostID  string `json:"post_id"`
	Content string `json:"content"`
}

func CommentCreation(w http.ResponseWriter, r *http.Request, uid int) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}
	var comment commentInfo
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if !PostComment(comment, uid) {
		w.WriteHeader(http.StatusInternalServerError)
		auth.JsRespond(w, "Comment creation failed", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	auth.JsRespond(w, "Comment posted successfully", http.StatusOK)
}

func PostComment(comment commentInfo, uid int) bool {
	query := `
INSERT
	INTO comments
	(post_id, uid, content)
VALUES
	(?, ?, ?)`
	_, err := helpers.DataBase.Exec(query,
		comment.PostID,
		uid,
		comment.Content)
	if err != nil {
		helpers.ErrorLog.Println("Database insertion error:", err)
		return false
	}
	return true
}
