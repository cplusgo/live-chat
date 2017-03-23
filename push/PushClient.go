package push

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"log"
	"github.com/cplusgo/live-chat/protocols"
)

type PushClient struct {
	wsConn       *websocket.Conn
	writeChannel chan []byte
}

func NewPushClient(conn *websocket.Conn) *PushClient {
	client := &PushClient{}
	return client
}

func (this *PushClient) ReadMessage() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("客户端主动断开连接")
		}
	}()
	for {
		_, originData, err := this.wsConn.ReadMessage()
		if err != nil {
			break
		}
		var message protocols.BaseMessageVo
		json.Unmarshal(originData, &message)
		switch(message.ProtocolId) {
		case protocols.MESSAGE_BROADCAST_PID:
			pushClientManager.broadcastChannel <- &message
		case protocols.REGISTER_PUSH_SERVER_PID:
			pushClientManager.addClientChannel <- this
		case protocols.UNREGISTER_PUSH_SERVRE_PID:
			pushClientManager.deleteClientChannel <- this
		}
	}
}

func (this *PushClient) waitMessage() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("客户端主动断开连接")
		}
	}()
	for {
		select {
		case message := <-this.writeChannel:
			this.wsConn.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func (this *PushClient) register() {

}
