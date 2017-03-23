package protocols

type BaseMessage struct {
	ProtocolId int `json:"protocol_id"`
	Body       string `json:"body"`
}
