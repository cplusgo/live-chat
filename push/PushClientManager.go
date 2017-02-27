package push

type PushClientManager struct {
	clients             map[*PushClient]*PushClient
	addClientChannel    chan *PushClient
	deleteClientChannel chan *PushClient
	broadcastChannel    chan *PushMessage
}

func NewPushClientManager() *PushClientManager {
	addChannel := make(chan *PushClient)
	deleteChannel := make(chan *PushClient)
	broadcastChannel := make(chan *PushMessage)
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

func (this *PushClientManager) broadcast(message *PushMessage) {
	for _, client := range this.clients {
		client.writeChannel <- message
	}
}

var pushClientManager *PushClientManager = NewPushClientManager()