package server

type ChatMessage struct {
	RoomId      string
	IsBroadcast bool
	Body        map[string]interface{}
}
