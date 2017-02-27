package push

type PushMessage struct {
	ProtocolId int64 `json:"protocolId"`
	Body       string `json:"body"`
	From       *PushClient
}
