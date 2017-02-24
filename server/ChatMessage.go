package server

type ChatMessage struct {
	ProtocolId  int64 `json:"protocolId"`
	RoomId      int64 `json:"roomId"`
	Data        string `json:"data"`
	isBroadcast bool
	body        map[string]interface{}
	originData  []byte
}

func NewChatMessage(roomId int64, body map[string]interface{}, originData []byte) *ChatMessage {
	message := &ChatMessage{
		RoomId:     roomId,
		body:       body,
		originData: originData,
	}
	return message
}
