package ws

import (
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1280,
	WriteBufferSize: 1280,
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	connection, err := upgrader.Upgrade(w, r, nil)
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if err = conn.WriteMessage(messageType, message); err != nil {
			return err
		}
	}
}
