package protocols

//进入房间协议号
const ENTER_ROOM_PID = 100
//房间消息协议号
const CHAT_MESSAGE_PID = 200
//消息广播
const MESSAGE_BROADCAST_PID = 201


/**以下是控制协议**/

//注册到中央状态服务器
const REGISTER_STATUS_SERVER_PID = 1001
//注册到中央消息服务器
const REGISTER_PUSH_SERVER_PID = 1002
//杀死某台弹幕服务器
const UNREGISTER_PUSH_SERVRE_PID = 1003
//汇报弹幕服务器的状态
const REPORT_STATUS_PID = 1004


