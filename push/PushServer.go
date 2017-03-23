package push

import (
	"github.com/gorilla/websocket"
	"net/http"
	"flag"
	"log"
	"github.com/cplusgo/go-library"
)

type PushServer struct {
	upgrader websocket.Upgrader
}

func NewPushServer() *PushServer {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	pushServer := &PushServer{upgrader:upgrader}
	return pushServer
}

func (this *PushServer) Start() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/websocket", this.accept)
	var addr = flag.String("addr", "localhost:8081", "http service address")
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func (this *PushServer) accept(w http.ResponseWriter, r *http.Request) {
	conn, err := this.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("PushServer.accept:", err)
		return
	}
	wsHandler := NewPushClient(conn)
	wsHandler.waitMessage()
	go wsHandler.ReadMessage()
}