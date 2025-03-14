package requests

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

func GetPosts() []Post {
	return []Post{}
}
