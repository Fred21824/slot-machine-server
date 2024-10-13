package websocket

import (
	"net/http"

	"slot-machine-server/internal/logger"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust this for production
	},
}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Error("Failed to upgrade to WebSocket", zap.Error(err))
		return
	}
	defer conn.Close()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			logger.Error("Error reading WebSocket message", zap.Error(err))
			return
		}

		if err := conn.WriteMessage(messageType, p); err != nil {
			logger.Error("Error writing WebSocket message", zap.Error(err))
			return
		}
	}
}
