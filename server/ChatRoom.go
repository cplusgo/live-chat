package server

import (
	"github.com/cplusgo/live-chat/protocols"
	"encoding/json"
	"log"
)

type ChatRoom struct {
	roomId           int
	clients          map[*ChatClient]*ChatClient
	broadcastChannel chan *protocols.ChatMessageVo
}

func NewChatRoom(roomId int) *ChatRoom {
	clients := make(map[*ChatClient]*ChatClient)
	broadcastChan := make(chan *protocols.ChatMessageVo)
	chatroom := &ChatRoom{
		roomId:           roomId,
		clients:          clients,
		broadcastChannel: broadcastChan,
	}
	go chatroom.run()
	return chatroom
}

func (this *ChatRoom) add(client *ChatClient) {
	this.clients[client] = client
}

func (this *ChatRoom) remove(client *ChatClient) {
	if _, ok := this.clients[client]; ok {
		delete(this.clients, client)
	}
}

func (this *ChatRoom) run() {
	for {
		select {
		case message := <-this.broadcastChannel:
			this.broadcastMessage(message)
		}
	}
}

func (this *ChatRoom) broadcastMessage(message *protocols.ChatMessageVo) {
	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err.Error())
		return
	}
	for _, client := range this.clients {
		client.writeChannel <- data
	}
}
