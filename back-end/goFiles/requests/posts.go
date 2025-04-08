package requests

import (
	"fmt"

	helpers "RTF/back-end"
)

type Post struct {
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	CreationTime string   `json:"creation_time"`
	Categories   []string `json:"categories"`
}

type Categories struct {
	ID   int
	Name string
}

// get from db
// CREATE TABLE IF NOT EXISTS `posts` (
//
//	post_id INTEGER PRIMARY KEY AUTOINCREMENT,
//	uid INTEGER NOT NULL,
//	title TEXT NOT NULL,
//	content TEXT NOT NULL,
//	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
//	FOREIGN KEY (uid) REFERENCES users(id) ON DELETE CASCADE
//
// );
func GetPosts() ([]Post, error) {
	rows, err := helpers.DataBase.Query(`
		SELECT title, content, created_at, uid
		FROM posts
	`)
	if err != nil {
		fmt.Println("Error getting posts: ", err)
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Title, &post.Content, &post.CreationTime, &post.Author)
		if err != nil {
			fmt.Println("Error scanning posts: ", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
