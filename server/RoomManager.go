package server

import "errors"

type RoomManager struct {
	rooms map[int64]*ChatRoom
	deleteClientChannel chan *ChatClient
	addClientChannel    chan *ChatClient
}

func NewRoomManager() *RoomManager {
	deleteChan := make(chan *ChatClient)
	addChan := make(chan *ChatClient)
	rooms := make(map[int64]*ChatRoom)
	manager := &RoomManager{
		deleteClientChannel:deleteChan,
		addClientChannel:addChan,
		rooms:rooms,
	}
	go manager.Run()
	return manager
}

func (this *RoomManager) Run()  {
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
		this.rooms[client.roomId] = NewChatRoom(client.roomId)
	}
	this.rooms[client.roomId].add(client)
}

func (this *RoomManager) sendMessage(roomId int64, message *ChatMessage)  {
	if room, ok := this.rooms[roomId]; ok {
		room.broadcastChannel <- message
	}
}

func (this *RoomManager) removeClient(client *ChatClient)  {
	if room, ok := this.rooms[client.roomId]; ok {
		room.remove(client)
	}
}

func (this *RoomManager) GetRoom(roomId int64) (*ChatRoom, error)  {
	if _, ok := this.rooms[roomId]; ok {
		return this.rooms[roomId], nil
	}
	return nil, errors.New("chat room not exist")
}

var roomManager *RoomManager = NewRoomManager()