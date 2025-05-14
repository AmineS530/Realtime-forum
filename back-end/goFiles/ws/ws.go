package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	// "RTF/back-end/goFiles/dms"
	helpers "RTF/back-end"
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

type update struct {
	Sender string `json:"sender"`
	Type   string `json:"type"`
	Uname  string `json:"username"`
	Online bool   `json:"online"`
}

type message struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

var (
	sockets = helpers.Sockets
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
		if string(msg)[:7] == "typing:" {
			m := string(msg)[7:]
			fmt.Println(m, "good")
			conn, exist := sockets[m]
			if !exist || conn == nil {
				continue
			}

			err := conn.WriteMessage(websocket.TextMessage, []byte(`{"sender":"internal","type":"typing","username":"`+uName+`"}`))
			fmt.Println(err)
			continue
		} else if string(msg)[:11] == "stoptyping:" {
			m := string(msg)[11:]
			fmt.Println(m, "good")
			conn, exist := sockets[m]
			if !exist || conn == nil {
				continue
			}

			err := conn.WriteMessage(websocket.TextMessage, []byte(`{"sender":"internal","type":"stoptyping","username":"`+uName+`"}`))
			fmt.Println(err)
			continue
		}
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
			status_response = `{"sender":"system","message":"failed to send message"}`
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
	} else {
		for _, v := range sockets {
			if err := v.WriteJSON(update{"internal", "toggle", uName, true}); err != nil {
				log.Println("azer azer azer azer", err)
			}
		}
	}
	sockets[uName] = conn
	mutex.Unlock()
}

func deleteConnFromMap(uName string) {
	mutex.Lock()
	delete(sockets, uName)
	for _, v := range sockets {
		if err := v.WriteJSON(update{"internal", "toggle", uName, false}); err != nil {
			log.Println("azer azer azer azer", err)
		}
	}
	mutex.Unlock()
}

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
		err = errors.New("failed to store message in db with error: " + err.Error())
		sockets[m.Sender].WriteJSON(message{"system", "", err.Error()})
		return err
	}
	responseData, err := json.Marshal(m)
	if err != nil {
		log.Println("Error marshaling response:", err)
		return err
	}
	conn, exist := sockets[m.Sender]
	if !exist || conn == nil {
		log.Printf("User %s not found or not connected\n", m.Receiver)
		return fmt.Errorf("user not found or not connected")
	}

	err = conn.WriteMessage(websocket.TextMessage, responseData)
	if err != nil {
		log.Println(err)
		return errors.New("failed to send message to receiver with error: " + err.Error())
	}

	conn, exist = sockets[m.Receiver]
	if !exist || conn == nil {
		return nil
	}

	conn.WriteMessage(websocket.TextMessage, responseData)
	return nil
}
