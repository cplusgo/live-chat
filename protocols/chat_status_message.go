package protocols

/*
搜集聊天服务器中的每个房间的用户数
*/
type ChatStatusMessageVo struct {
	RoomStatusList []ChatRoomStatusVo `json:"room_status_list"`
}

type ChatRoomStatusVo struct {
	roomId    int        `json:"room_id"`
	clientNum int        `json:"client_num"`
}
