package channel

type Channel struct {
	id                 int
	clients            map[*Consumer]*Consumer
	addConsumerChan    chan *Consumer
	removeConsumerChan chan *Consumer
}

func NewChannel(id int) *Channel {
	channel := &Channel{

	}
	return channel
}
