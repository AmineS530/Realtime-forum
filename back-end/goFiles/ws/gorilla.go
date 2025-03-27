package ws

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool {
// 		return true // ⚠ Change this for security in production
// 	},
// }

// var (
// 	clients   = make(map[*websocket.Conn]bool)
// 	broadcast = make(chan []byte)
// )

// // WebSocket handler
// func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		fmt.Println("WebSocket upgrade failed:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	clients[conn] = true
// 	fmt.Println("New client connected")

// 	// Listen for messages from this client
// 	for {
// 		_, msg, err := conn.ReadMessage()
// 		if err != nil {
// 			fmt.Println("Client disconnected")
// 			delete(clients, conn)
// 			break
// 		}
// 		fmt.Printf("Received: %s\n", msg)
// 		broadcast <- msg // Send to all clients
// 	}
// }

// // Broadcast messages to all connected clients
// func Broadcaster() {
// 	for {
// 		msg := <-broadcast
// 		for client := range clients {
// 			err := client.WriteMessage(websocket.TextMessage, msg)
// 			if err != nil {
// 				client.Close()
// 				delete(clients, client)
// 			}
// 		}
// 	}
// }

// needed to force logout old user sessions
// func NotifyOldSessions(userID int) {
// 	message := fmt.Sprintf(`{"action": "logout", "message": "Logged out from another device"}`)
// 	websocketServer.BroadcastToUser(userID, message)
// }
