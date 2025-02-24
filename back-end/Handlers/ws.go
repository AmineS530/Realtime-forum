package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Define a WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow connections from any origin
		return true
	},
}

type Transmission struct {
	Type    string          `json:"type"`
	Content json.RawMessage `json:"content"`
	Sender  string          `json:"-"`
}

type msg struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

var (
	sockets = make(map[string]*websocket.Conn)
	mutex   sync.Mutex
)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	uName := getUname(r)
	fmt.Printf("New connection: %s\n", uName)
	mutex.Lock()
	if conn, exists := sockets[uName]; exists {
		log.Printf("User %s already connected\n", uName)
		conn.Close()
	}
	sockets[uName] = conn
	defer delete(sockets, uName)
	mutex.Unlock()
	for {
		// Read message from the WebSocket connection
		azer, msg, err := conn.ReadMessage()
		if err != nil || azer != websocket.TextMessage {
			log.Println(err)
			return
		}
		fmt.Println("ws", azer, string(msg), err)

		// Print the received message
		fmt.Printf("Received: [%s]-->%s\n", uName, msg)

		var request Transmission
		err = json.Unmarshal(msg, &request)
		if err != nil {
			log.Println("Error parsing JSON:", err)
			continue
		}
		request.Sender = uName

		// Print the parsed message
		fmt.Printf("Parsed Transmission: %+v\n", request)

		// Respond back with a JSON message
		err = request.handle()
		var status_response string
		if err != nil {
			log.Println("Error handling request:", err)
			status_response = `{"type":"status","content":"failed"}`
		} else {
			status_response = `{"type":"status","content":"sucess"}`
		}

		// Send the JSON response back to the client
		fmt.Printf("Sending response: %+v\n", status_response)
		err = conn.WriteMessage(websocket.TextMessage, []byte(status_response))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

var i rune = '0'

func getUname(r *http.Request) string {
	uname := r.URL.Query().Get("uname")
	if uname == "" {
		uname = "guest" + string(i)
		i++
	}
	return uname
}

func (m *msg) send() error {
	conn, exist := sockets[m.Receiver]
	if !exist || conn == nil {
		log.Printf("User %s not found or not connected\n", m.Receiver)
		return fmt.Errorf("user not found or not connected")
	}

	msg, err := json.Marshal(m)
	if err != nil {
		log.Println("Error marshaling response:", err)
		return err
	}
	responseData, err := json.Marshal(Transmission{Type: "message", Content: []byte(msg)})
	if err != nil {
		log.Println("Error marshaling response:", err)
		return err
	}

	err = conn.WriteMessage(websocket.TextMessage, responseData)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (m *Transmission) handle() error {
	var message msg
	err := json.Unmarshal(m.Content, &message)
	if err != nil {
		log.Println("Error parsing JSON:", err)
		return err
	}
	message.Sender = m.Sender
	fmt.Println("sending", message)
	return message.send()
}

func (t *Transmission) new(Type string, content any) error {
	msg, err := json.Marshal(content)
	if err != nil {
		return err
	}
	t.Type = Type
	t.Content = msg
	return nil
}
