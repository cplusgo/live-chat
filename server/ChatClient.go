package server

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"log"
	"github.com/cplusgo/live-chat/protocols"
	"fmt"
)

type ChatClient struct {
	wsConn       *websocket.Conn
	roomId       int
	writeChannel chan *protocols.ChatMessage
	stopChannel  chan bool
}

func NewChatClient(conn *websocket.Conn) *ChatClient {
	writeChannel := make(chan *protocols.ChatMessage)
	stopChannel := make(chan bool)
	client := &ChatClient{wsConn: conn, writeChannel: writeChannel, stopChannel: stopChannel}
	go client.readMessage()
	return client
}

func (this *ChatClient) close() {
	if this.roomId != 0 {
		this.wsConn.Close()
		this.stopChannel <- true
		roomManager.deleteClientChannel <- this
	}
}

func (this *ChatClient) readMessage() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("客户端主动断开连接")
			this.close()
		}
	}()
	for {
		_, bytes, err := this.wsConn.ReadMessage()
		if err != nil {
			this.close()
			break
		}
		log.Println(string(bytes))
		var message protocols.BaseMessage
		err = json.Unmarshal(bytes, &message)
		if err != nil {
			log.Println(err)
		} else {
			protocolId := message.ProtocolId
			switch protocolId {
			case protocols.ENTER_ROOM_ID:
				this.enterRoom(&message)
			case protocols.CHAT_MESSAGE_ID:
				this.broadcastInRoom(&message)
			}
		}
	}
}

func (this *ChatClient) broadcastInRoom(message *protocols.BaseMessage) {
	var chatMessage protocols.ChatMessage
	json.Unmarshal([]byte(message.Body), &chatMessage)
	roomManager.sendMessage(this.roomId, &chatMessage)
}

func (this *ChatClient) enterRoom(message *protocols.BaseMessage) {
	var enterRoomMessage protocols.EnterRoomMessage
	json.Unmarshal([]byte(message.Body), &enterRoomMessage)
	this.roomId = enterRoomMessage.RoomId
	roomManager.addClientChannel <- this
}

func (this *ChatClient) writeMessage(message []byte) {
	if !this.isWritable() {
		return
	}

	err := this.wsConn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		this.close()
	}
}

func (this *ChatClient) isWritable() bool {
	return true
}

func (this *ChatClient) Try() {
	for {
		select {
		case message := <-this.writeChannel:
			fmt.Println(message.Content)
			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err.Error())
				return
			}
			this.writeMessage(data)
		case <-this.stopChannel:
			return
		}
	}
}

func (this *ChatClient) Catch(err interface{}) {
	this.close()
}
