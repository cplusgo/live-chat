package server

import (
	"github.com/gorilla/websocket"
	"log"
)

type WsHandler struct {
	WsConn *websocket.Conn
}

func (this *WsHandler) Try() {
	for {
		mt, message, err := this.WsConn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = this.WsConn.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
	this.WsConn.Close()
}

func (this *WsHandler) Catch(err interface{}) {
	this.WsConn.Close()
}
