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
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if err = conn.WriteMessage(messageType, p); err != nil {
			return err
		}
	}
}
