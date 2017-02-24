package server

type ChatRoom struct {
	roomId              int64
	clients             map[*ChatClient]*ChatClient
	broadcastChannel    chan *ChatMessage

}

func NewChatRoom(roomId int64) *ChatRoom {
	clients := make(map[*ChatClient]*ChatClient)
	broadcastChan := make(chan *ChatMessage)
	chatroom := &ChatRoom{
		roomId:              roomId,
		clients:             clients,
		broadcastChannel:    broadcastChan,
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

func (this *ChatRoom) broadcastMessage(message *ChatMessage) {
	for _, client := range this.clients {
		client.writeChannel <- message
	}
}
