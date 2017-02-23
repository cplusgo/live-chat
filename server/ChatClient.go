package server

type ChatClient struct {
	RoomId string
}

func (this *ChatClient) IsWritable() bool {
	return true
}

func (this *ChatClient) SendMessage(message *ChatMessage) {
	if this.IsWritable() {

	}
}