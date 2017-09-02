package channel

import (
	"github.com/gorilla/websocket"
)

type Consumer struct {
	wsConn       *websocket.Conn
	roomId       int
	writeChannel chan []byte
	stopChannel  chan bool
}
