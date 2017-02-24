package server

import "errors"

type RoomManager struct {
	rooms map[string]*ChatRoom
}

func NewRoomManager() *RoomManager {
	manager := &RoomManager{}
	return manager
}

func (this *RoomManager) AddClient(client *ChatClient) {
	if _, ok := this.rooms[client.RoomId]; !ok {
		this.rooms[client.RoomId] = NewChatRoom(client.RoomId)
	}
	this.rooms[client.RoomId].Add(client)
}

func (this *RoomManager) RemoveClient(client *ChatClient)  {
	if _, ok := this.rooms[client.RoomId]; ok {
		this.rooms[client.RoomId].Remove(client)
	}
}

func (this *RoomManager) GetRoom(roomId string) (*ChatRoom, error)  {
	if _, ok := this.rooms[roomId]; ok {
		return this.rooms[roomId], nil
	}
	return nil, errors.New("chat room not exist")
}

var roomManager *RoomManager = NewRoomManager()