package server

import (
	"github.com/gorilla/websocket"
	"log"
	"encoding/json"
)

type ChatClient struct {
	wsConn       *websocket.Conn
	RoomId       string
	WriteChannel chan []byte
	ReadChannel  chan []byte
}

func (this *ChatClient) Close() {
	this.wsConn.Close()
}

func (this *ChatClient) Try() {
	for {
		err := this.ReadMessage()
		if err != nil {
			break
		}
	}
}

func (this *ChatClient) ReadMessage() error {
	_, message, err := this.wsConn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		this.Close()
		return err
	}
	this.Process(message)
	return nil
}

func (this *ChatClient) WriteMessage(message []byte) {
	err := this.wsConn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		this.Close()
	}
}

func (this *ChatClient) Process(bytes []byte) {
	var data interface{}
	json.Unmarshal(bytes, &data)
	body := data.(map[string]interface{})
	protocolId := 0
	if id, ok := body[PROTOCOL_ID]; ok {
		protocolId = id.(int)
	}
	switch protocolId {
	case P_LOGIN_ROOM:
		this.AddSelfInRoom(body)
	}
}

func (this *ChatClient) AddSelfInRoom(body map[string]interface{}) {
	roomManager.AddClient(this)
}

func (this *ChatClient) IsWritable() bool {
	return true
}

func (this *ChatClient) BroadcastInRoom() {

}

func (this *ChatClient) SendMessage(message *ChatMessage) {
	if this.IsWritable() {
		room, err := roomManager.GetRoom(this.RoomId)
		if err == nil {
			room.broadcastChannel <- message
		}
	}
}

func (this *ChatClient) Catch(err interface{}) {
	this.Close()
}
