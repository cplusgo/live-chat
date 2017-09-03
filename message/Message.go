package message

/*
 * 通用消息格式，服务器不关心Body中的数据，对服务器而言Body就是一串字节
 */
type Message struct {
	CommandId int
	Length    int
	Body      []byte
}
