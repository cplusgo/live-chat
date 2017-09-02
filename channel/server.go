package channel

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

type ChannelServer struct {
	upgrader websocket.Upgrader
}

func NewChannelServer() *ChannelServer {
	chatServer := &ChannelServer{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	return chatServer
}

func (this *ChannelServer) Start() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/websocket", this.accept)
	var addr = flag.String("addr", "localhost:8080", "http service address")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func (this *ChannelServer) accept(w http.ResponseWriter, r *http.Request) {
	conn, err := this.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ChannelServer.accept:", err)
		return
	}
	startChatClient(conn)
}
