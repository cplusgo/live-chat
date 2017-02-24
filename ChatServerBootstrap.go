package main

import "github.com/cplusgo/live-chat/server"

func main() {
	chatServer := server.NewChatServer()
	chatServer.Start()
}