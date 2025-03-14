package ws

import (
	"fmt"

	_ "github.com/gorilla/websocket"
)

// needed to force logout old user sessions
// func NotifyOldSessions(userID int) {
// 	message := fmt.Sprintf(`{"action": "logout", "message": "Logged out from another device"}`)
// 	websocketServer.BroadcastToUser(userID, message)
// }
