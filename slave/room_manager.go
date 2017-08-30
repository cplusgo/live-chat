package slave

import (
	"errors"
	"github.com/cplusgo/live-chat/protocols"
)

type RoomManager struct {
	rooms map[int]*ChatRoom
	deleteClientChannel chan *ChatClient
	addClientChannel    chan *ChatClient
}

func createRoomManager() *RoomManager {
	deleteChan := make(chan *ChatClient)
	addChan := make(chan *ChatClient)
	rooms := make(map[int]*ChatRoom)
	manager := &RoomManager{
		deleteClientChannel:deleteChan,
		addClientChannel:addChan,
		rooms:rooms,
	}
	go manager.run()
	return manager
}

func (this *RoomManager) run()  {
	for {
		select {
		case client := <-this.deleteClientChannel:
			this.removeClient(client)
		case client := <-this.addClientChannel:
			this.addClient(client)
		}
	}
}

func (this *RoomManager) addClient(client *ChatClient) {
	if _, ok := this.rooms[client.roomId]; !ok {
		this.rooms[client.roomId] = createChatRoom(client.roomId)
	}
	this.rooms[client.roomId].add(client)
}

func (this *RoomManager) sendMessage(roomId int, message *protocols.ChatMessageVo)  {
	if room, ok := this.rooms[roomId]; ok {
		room.broadcastChannel <- message
	}
}

func (this *RoomManager) removeClient(client *ChatClient)  {
	if room, ok := this.rooms[client.roomId]; ok {
		room.remove(client)
	}
}

func (this *RoomManager) getRoomById(roomId int) (*ChatRoom, error)  {
	if _, ok := this.rooms[roomId]; ok {
		return this.rooms[roomId], nil
	}
	return nil, errors.New("chat room not exist")
}

var roomManager *RoomManager = createRoomManager()