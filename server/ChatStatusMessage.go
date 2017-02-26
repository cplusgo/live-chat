package server

/***
搜集聊天服务器中的每个房间的用户数
***/
type ChatStatusMessage struct {
	RoomStatusList []ChatRoomStatus
}


type ChatRoomStatus struct {
	roomId int64
	clientNum int32
}