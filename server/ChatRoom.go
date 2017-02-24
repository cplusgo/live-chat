package server

type ChatRoom struct {
	roomId         string
	clients        map[*ChatClient]*ChatClient
	broadcastChannel *ChatMessage
}

func NewChatRoom(roomId string) *ChatRoom {
	clients := make(map[*ChatClient]*ChatClient)
	chatroom := &ChatRoom{roomId: roomId, clients: clients}
	return chatroom
}

func (this *ChatRoom) Add(client *ChatClient) {
	this.clients[client] = client
}

func (this *ChatRoom) Remove(client *ChatClient) {
	if _, ok := this.clients[client]; ok {
		delete(this.clients, client)
	}
}

func (this *ChatRoom) Run()  {
	for {
		select {
		case message := <-this.broadcastChannel:
			this.BroadcastMessage(message)
		}
	}
}

func (this *ChatRoom) BroadcastMessage(message *ChatMessage) {
	for _, client := range this.clients {
		client.SendMessage(message)
	}
}
