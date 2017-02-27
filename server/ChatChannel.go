package server

type ChatChannel struct {
	roomId              int64
	clients             map[*ChatClient]*ChatClient
	broadcastChannel    chan *ChatMessage

}

func NewChatChannel(roomId int64) *ChatChannel {
	clients := make(map[*ChatClient]*ChatClient)
	broadcastChan := make(chan *ChatMessage)
	chatroom := &ChatChannel{
		roomId:              roomId,
		clients:             clients,
		broadcastChannel:    broadcastChan,
	}
	go chatroom.run()
	return chatroom
}

func (this *ChatChannel) add(client *ChatClient) {
	this.clients[client] = client
}

func (this *ChatChannel) remove(client *ChatClient) {
	if _, ok := this.clients[client]; ok {
		delete(this.clients, client)
	}
}

func (this *ChatChannel) run() {
	for {
		select {
		case message := <-this.broadcastChannel:
			this.broadcastMessage(message)
		}
	}
}

func (this *ChatChannel) broadcastMessage(message *ChatMessage) {
	for _, client := range this.clients {
		client.writeChannel <- message
	}
}
