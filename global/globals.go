package global

import "github.com/gorilla/websocket"

var Sockets = make(map[string]*websocket.Conn)
