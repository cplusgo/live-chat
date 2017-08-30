package protocols

type BaseMessageVo struct {
	ProtocolId int `json:"protocol_id"`
	Body       string `json:"body"`
}
