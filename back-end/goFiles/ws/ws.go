package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	// "RTF/back-end/goFiles/dms"
	jwt "RTF/back-end/goFiles/JWT"
	"RTF/back-end/goFiles/auth"
	"RTF/back-end/goFiles/dms"

	"github.com/gorilla/websocket"
)

// Define a WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

var (
	sockets = make(map[string]*websocket.Conn)
	mutex   sync.Mutex
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	uName := getUname(r)
	fmt.Printf("New connection: %s\n", uName)

	addConnToMap(uName, conn)
	defer deleteConnFromMap(uName)

	for {
		// Read message from the WebSocket connection
		azer, msg, err := conn.ReadMessage()
		if err != nil || azer != websocket.TextMessage {
			log.Println(err)
			return
		}
		fmt.Println("ws", azer, string(msg), err)

		var request message
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			continue
		}
		request.Sender = uName

		// Respond back with a JSON message
		err = request.send()
		var status_response string
		fmt.Println(sockets)
		if err != nil {
			log.Println("Error handling request:", err)
			status_response = `{"sender":"system","content":"failed to send message"}`
		} else {
			fmt.Printf("Sending response: %+v\n", status_response)
			err = conn.WriteMessage(websocket.TextMessage, []byte(status_response))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func addConnToMap(uName string, conn *websocket.Conn) {
	mutex.Lock()
	if conn, exists := sockets[uName]; exists {
		log.Printf("User %s already connected\n", uName)
		conn.Close()
	}
	sockets[uName] = conn
	mutex.Unlock()
}

func deleteConnFromMap(uName string) {
	mutex.Lock()
	delete(sockets, uName)
	mutex.Unlock()
}

// var i rune = '0'

// TODO JWT
func getUname(r *http.Request) string {
	payload := r.Context().Value(auth.UserContextKey)
	data, ok := payload.(*jwt.JwtPayload)
	if ok {
		return data.Username
	} else {
		return ""
	}
}

func (m *message) send() error {
	err := dms.AddDm(m.Sender, m.Receiver, m.Message)
	if err != nil {
		return errors.New("failed to store message in db with error: " + err.Error())
	}
	conn, exist := sockets[m.Receiver]
	if !exist || conn == nil {
		log.Printf("User %s not found or not connected\n", m.Receiver)
		return fmt.Errorf("user not found or not connected")
	}

	responseData, err := json.Marshal(m)
	if err != nil {
		log.Println("Error marshaling response:", err)
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, responseData)
	if err != nil {
		log.Println(err)
		return errors.New("failed to send message to receiver with error: " + err.Error())
	}
	return nil
}
