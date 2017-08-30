package slave

import (
	"github.com/gorilla/websocket"
	"encoding/json"
	"log"
	"github.com/cplusgo/live-chat/protocols"
)

type ChatClient struct {
	wsConn       *websocket.Conn
	roomId       int
	writeChannel chan []byte
	stopChannel  chan bool
}

func startChatClient(conn *websocket.Conn) *ChatClient {
	writeChannel := make(chan []byte)
	stopChannel := make(chan bool)
	client := &ChatClient{wsConn: conn, writeChannel: writeChannel, stopChannel: stopChannel}
	go client.onWaitMessageIn()
	client.onWaitMessageOut()
	return client
}

func (this *ChatClient) close() {
	log.Println("客户端主动断开连接")
	if this.roomId != 0 {
		this.wsConn.Close()
		this.stopChannel <- true
		roomManager.deleteClientChannel <- this
		_, isClose := <-this.stopChannel
		if !isClose {
			close(this.stopChannel)
		}
	}
}

func (this *ChatClient) onWaitMessageIn() {
	defer func() {
		if err := recover(); err != nil {
			this.close()
		}
	}()
	for {
		_, bytes, err := this.wsConn.ReadMessage()
		if err != nil {
			log.Println(this)
			this.close()
			return
		}
		log.Println(string(bytes))
		var message protocols.BaseMessageVo
		err = json.Unmarshal(bytes, &message)
		if err != nil {
			log.Println(err)
		} else {
			protocolId := message.ProtocolId
			switch protocolId {
			case protocols.ENTER_ROOM_PID:
				this.onEnterRoom(&message)
			case protocols.CHAT_MESSAGE_PID:
				this.onMessageReceived(&message)
			}
		}
	}
}

func (this *ChatClient) onMessageReceived(message *protocols.BaseMessageVo) {
	var chatMessage protocols.ChatMessageVo
	json.Unmarshal([]byte(message.Body), &chatMessage)
	roomManager.sendMessage(this.roomId, &chatMessage)
}

func (this *ChatClient) onEnterRoom(message *protocols.BaseMessageVo) {
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

func (this *ChatClient) onWaitMessageOut() {
	defer func() {
		if err := recover(); err != nil {
			this.close()
		}
	}()
	for {
		select {
		case message := <-this.writeChannel:
			this.writeMessage(message)
		case <-this.stopChannel:
			return
		}
	}
}
