package protocols

import "github.com/cplusgo/live-chat/push"

type PushMessageVo struct {
	Data []byte `json:"data"`
	From *push.PushClient `json:"from"`
}
