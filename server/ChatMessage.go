package server

type ChatMessage struct {
	roomId      int64
	isBroadcast bool
	body        map[string]interface{}
	originData []byte
}

func NewChatMessage(roomId int64, body map[string]interface{}, originData []byte) *ChatMessage {
	message := &ChatMessage{
		roomId:roomId,
		body:body,
		originData:originData,
	}
	return message
}