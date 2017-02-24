package server

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"log"
)

type ChatClient struct {
	wsConn       *websocket.Conn
	roomId       int64
	writeChannel chan *ChatMessage
}

func NewChatClient(conn *websocket.Conn) *ChatClient {
	writeChannel := make(chan *ChatMessage)
	client := &ChatClient{wsConn: conn, writeChannel: writeChannel}
	go client.readMessage()
	return client
}

func (this *ChatClient) close() {
	if this.roomId != 0 {
		this.wsConn.Close()
		roomManager.deleteClientChannel <- this
	}
}

func (this *ChatClient) readMessage() {
	for {
		_, bytes, err := this.wsConn.ReadMessage()
		if err != nil {
			this.close()
			break
		}
		log.Println(string(bytes))
		var message ChatMessage
		err = json.Unmarshal(bytes, &message)
		if err != nil {
			log.Println(err)
		} else {
			protocolId := message.ProtocolId
			message.originData = bytes
			switch protocolId {
			case P_LOGIN_ROOM:
				this.registerInRoom(&message)
			case P_NORMAL_MSG:
				this.broadcastInRoom(&message)
			}
		}
	}
}

func (this *ChatClient) broadcastInRoom(message *ChatMessage) {
	roomManager.sendMessage(this.roomId, message)
}

func (this *ChatClient) registerInRoom(message *ChatMessage) {
	roomManager.addClientChannel <- this
}

func (this *ChatClient) writeMessage(message *ChatMessage) {
	err := this.wsConn.WriteMessage(websocket.TextMessage, message.originData)
	if err != nil {
		this.close()
	}
}

func (this *ChatClient) IsWritable() bool {
	return true
}

func (this *ChatClient) Try() {
	for {
		select {
		case message := <-this.writeChannel:
			this.writeMessage(message)
		}
	}
}

func (this *ChatClient) Catch(err interface{}) {
	this.close()
}
