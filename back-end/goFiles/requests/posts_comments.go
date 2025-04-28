package requests

import (
	"fmt"
	"strconv"
	"strings"

	"RTF/global"
)

type Post struct {
	Pid          int      `json:"pid"`
	Title        string   `json:"title"`
	Content      string   `json:"content"`
	Author       string   `json:"author"`
	CreationTime string   `json:"creation_time"`
	Categories   []string `json:"categories"`
}

type Comment struct {
	Pid          int    `json:"pid"`
	Author       string `json:"author"`
	Content      string `json:"content"`
	CreationTime string `json:"creation_time"`
}

func GetPosts(soffset string) ([]Post, error) {
	offset, err := strconv.Atoi(soffset)
	if err != nil {
		global.ErrorLog.Println("Error converting pid to int: ", err)
		return nil, err
	}
	rows, err := global.DataBase.Query(`
	SELECT 
   		p.post_id AS pid, 
    	p.title, 
    	p.content, 
    	p.categories, 
    	p.created_at, 
    	u.username AS author
	FROM
    	posts p
	JOIN 
    	users u ON p.uid = u.id
	LIMIT 3 OFFSET ?
	`, offset)
	if err != nil {
		fmt.Println("Error getting posts: ", err)
		return nil, err
	}
	defer rows.Close()
	var categories string
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.Pid, &post.Title, &post.Content, &categories, &post.CreationTime, &post.Author)
		if err != nil {
			fmt.Println("Error scanning posts: ", err)
			return nil, err
		}
		post.Categories = strings.Split(categories, ", ")
		posts = append(posts, post)
	}
	return posts, nil
}

func GetComments(pid string) ([]Comment, error) {
	iPid, err := strconv.Atoi(pid)
	if err != nil {
		global.ErrorLog.Println("Error converting pid to int: ", err)
		return nil, err
	}
	rows, err := global.DataBase.Query(`
	SELECT 
    	u.username AS author,
    	c.content,
    	c.created_at
	FROM 
    	comments c
	JOIN 
    	users u ON c.uid = u.id
	WHERE
		c.post_id = ?
	ORDER BY
		c.created_at DESC
	`, iPid)
	if err != nil {
		fmt.Println("Error getting comments: ", err)
		return nil, err
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var comment Comment
		err := rows.Scan(&comment.Author, &comment.Content, &comment.CreationTime)
		if err != nil {
			fmt.Println("Error scanning comments: ", err)
			return nil, err
		}
		comment.Pid = iPid
		comments = append(comments, comment)
	}
	return comments, nil
}
