package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	helpers "RTF/back-end"
	jwt "RTF/back-end/goFiles/JWT"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/requests"
	"RTF/back-end/goFiles/ws"
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/api/v1/ws", ws.HandleConnections)
	mux.HandleFunc("/api/v1/get/{type}", auth.AuthMiddleware(dumbjson))
	mux.HandleFunc("/api/profile", auth.AuthMiddleware(ProfileHandler))
	mux.HandleFunc("/api/ws", ws.HandleWebSocket)
	ProtectedStatic(mux, "/front-end/styles/", "./front-end/styles")
	ProtectedStatic(mux, "/front-end/scripts/", "./front-end/scripts")
	ProtectedStatic(mux, "/front-end/images/", "./front-end/images")
	// mux.Handle("/front-end/scripts/", http.StripPrefix("/front-end/scripts/", http.FileServer(http.Dir("./front-end/scripts"))))
	// mux.Handle("/front-end/images/", http.StripPrefix("/front-end/images/", http.FileServer(http.Dir("./front-end/images"))))
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
		usernames, _ := getUserNames()
		jsonData, _ := json.Marshal(usernames)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	case "dmhistory":
		target := r.Header.Get("target")
		tok, _ := auth.ExtractJWT(r)
		payload, _ := jwt.JWTVerify(tok)
		username := payload.Username
		dms, _ := getdmHistory(username, target)
		jsonData, _ := json.Marshal(dms)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
		fmt.Println(target, username, dms)
	default:
		helpers.ErrorPagehandler(w, http.StatusNotFound)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if helpers.Err != nil {
		helpers.ErrorPagehandler(w, http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/api/") {
		helpers.ErrorPagehandler(w, http.StatusNotFound)
		return
	}
	if r.URL.Path == "/" || r.URL.Path == "/profile" {
		if err := helpers.HtmlTemplates.ExecuteTemplate(w, "index.html", nil); err != nil {
			fmt.Println("Error executing template: ", err.Error())
			helpers.ErrorPagehandler(w, http.StatusInternalServerError)
			return
		}
	} else {
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

// TODO Their propper place
type Message struct {
	Sender  string `json:"sender"`
	Content string `json:"message"`
}

func getdmHistory(uname1, uname2 string) ([]Message, error) {
	rows, err := helpers.DataBase.Query(`
SELECT sender.username , d.message
FROM dms d
JOIN users sender ON d.sender_id = sender.id
JOIN users recipient ON d.recipient_id = recipient.id
WHERE (sender.username = ? AND recipient.username = ?)
   OR (sender.username = ? AND recipient.username = ?);
	`, uname1, uname2, uname2, uname1)
	if err != nil {
		fmt.Println("Error getting posts: ", err)
		return nil, err
	}
	defer rows.Close()
	var messages []Message
	for rows.Next() {
		var message Message
		err := rows.Scan(&message.Sender, &message.Content)
		if err != nil {
			fmt.Println("Error scanning posts: ", err)
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func getUserNames() ([]string, error) {
	rows, err := helpers.DataBase.Query("SELECT username FROM users")
	if err != nil {
		return nil, fmt.Errorf("could not execute query: %w", err)
	}
	defer rows.Close() // Make sure to close the rows when done

	var userNames []string

	// Iterate over the result set
	for rows.Next() {
		var username string
		if err := rows.Scan(&username); err != nil {
			return userNames, fmt.Errorf("could not scan row: %w", err)
		}
		userNames = append(userNames, username)
	}

	// Check if there was an error iterating over the rows
	if err := rows.Err(); err != nil {
		return userNames, fmt.Errorf("row iteration error: %w", err)
	}

	return userNames, nil
}
