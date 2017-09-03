package channel

import "github.com/cplusgo/live-chat/message"

type Channel struct {
	id                 int
	clients            map[*Consumer]*Consumer
	addConsumerChan    chan *Consumer
	removeConsumerChan chan *Consumer
	broadcastChan      chan *message.Message
}

func NewChannel(id int) *Channel {
	channel := &Channel{

	}
	return channel
}
