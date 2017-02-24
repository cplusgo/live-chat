package server

type ChatMessage struct {
	ProtocolId  int64 `json:"protocolId"`
	RoomId      int64 `json:"roomId"`
	Data        string `json:"data"`
	isBroadcast bool
	originData  []byte
}
