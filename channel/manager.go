package channel

import (
	"log"
	"fmt"
	"errors"
)

const (
	MAX_CHANNEL_LENGTH = 100
)

/*
 * 频道管理
 */
type ChannelManager struct {
	channels          map[int]*Channel
	addChannelChan    chan *Channel
	removeChannelChan chan *Channel
}

var channelManager *ChannelManager

func NewChannelManager() *ChannelManager {
	if channelManager == nil {
		channelManager = &ChannelManager{
			channels:          make(map[int]*Channel),
			addChannelChan:    make(chan *Channel, MAX_CHANNEL_LENGTH),
			removeChannelChan: make(chan *Channel, MAX_CHANNEL_LENGTH),
		}
		go channelManager.waitLoop()
	}

	return channelManager
}

func (this *ChannelManager) Add(channel *Channel) {
	this.addChannelChan <- channel
}

func (this *ChannelManager) Remove(channel *Channel) {
	this.removeChannelChan <- channel
}

func (this *ChannelManager) waitLoop() {
	for {
		select {
		case channel := <-this.addChannelChan:
			this.addImpl(channel)
		case channel := <-this.removeChannelChan:
			this.removeImpl(channel)
		}
	}
}

func (this *ChannelManager) addImpl(channel *Channel) error {
	if _, ok := this.channels[channel.id]; ok {
		msg := fmt.Sprintf("Channel %d already exists", channel.id)
		log.Println(msg)
		return errors.New(msg)
	}
	this.channels[channel.id] = channel
	return nil
}

func (this *ChannelManager) removeImpl(channel *Channel) error {
	if _, ok := this.channels[channel.id]; !ok {
		msg := fmt.Sprintf("The channel %d doesn't exist", channel.id)
		log.Println(msg)
		return errors.New(msg)
	}
	delete(this.channels, channel.id)
	return nil
}
