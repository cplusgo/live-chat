package server

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/cplusgo/go-library"
)

type ChatServer struct {
	upgrader websocket.Upgrader
}

func NewChatServer() *ChatServer {
	chatServer := &ChatServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	return chatServer
}

func (this *ChatServer) Start() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", this.Accept)
	var addr = flag.String("addr", "localhost:8080", "http service address")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func (this *ChatServer) Accept(w http.ResponseWriter, r *http.Request) {
	c, err := this.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	wsHandler := &WsHandler{wsConn: c}
	go_library.Run(wsHandler)
}
