package push

import (
	"github.com/gorilla/websocket"
	"encoding/json"
)

type PushClient struct {
	wsConn       *websocket.Conn
	writeChannel chan *PushMessage
}

func NewPushClient(conn *websocket.Conn) *PushClient {
	client := &PushClient{}
	return client
}

func (this *PushClient) Try() {
	for {
		_, originData, err := this.wsConn.ReadMessage()
		if err != nil {
			break
		}
		var message PushMessage
		json.Unmarshal(originData, &message)
		message.From = this
		switch(message.ProtocolId) {
		case MESSAGE_BROATCAST:
			pushClientManager.broadcastChannel <- &message
		case MESSAGE_REGISTER:
			pushClientManager.addClientChannel <- this
		case MESSAGE_KILL_ME:
			pushClientManager.deleteClientChannel <- this
		}
	}
}

func (this *PushClient) register()  {
	
}

func (this *PushClient) Catch(err interface{}) {

}
