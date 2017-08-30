package push

import "github.com/cplusgo/live-chat/protocols"

type PushClientManager struct {
	clients             map[*PushClient]*PushClient
	addClientChannel    chan *PushClient
	deleteClientChannel chan *PushClient
	broadcastChannel    chan *protocols.PushMessageVo
}

func NewPushClientManager() *PushClientManager {
	addChannel := make(chan *PushClient)
	deleteChannel := make(chan *PushClient)
	broadcastChannel := make(chan *protocols.PushMessageVo)
	clients := make(map[*PushClient]*PushClient)

	manager := &PushClientManager{
		clients:             clients,
		addClientChannel:    addChannel,
		deleteClientChannel: deleteChannel,
		broadcastChannel:    broadcastChannel,
	}
	return manager
}

func (this *PushClientManager) Run() {
	for {
		select {
		case client := <-this.addClientChannel:
			this.addClient(client)
		case client := <-this.deleteClientChannel:
			this.deleteClient(client)
		case message := <-this.broadcastChannel:
			this.broadcast(message)
		}
	}
}

func (this *PushClientManager) addClient(client *PushClient) {
	if _, ok := this.clients[client]; !ok {
		this.clients[client] = client
	}
}

func (this *PushClientManager) deleteClient(client *PushClient) {
	if _, ok := this.clients[client]; ok {
		delete(this.clients, client)
	}
}

func (this *PushClientManager) broadcast(message *protocols.PushMessageVo) {
	data := []byte(message.Data)
	for _, client := range this.clients {
		if client != message.From {
			client.writeChannel <- data
		}
	}
}

var pushClientManager *PushClientManager = NewPushClientManager()