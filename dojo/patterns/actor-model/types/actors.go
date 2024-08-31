package types

import "fmt"

type IActor interface {
	SendMessage(msg Message)
	ProcessMessages()
}

type Actor struct {
	Name string
	Messages []Message
}

type RemoteActor struct {}

func (a *Actor) SendMessage(msg Message) {
	a.Messages = append(a.Messages, msg)
}

func (a *Actor) ProcessMessages() {
	for _, m := range a.Messages {
		fmt.Printf("Actor %s received message: %s\n", a.Name, m.Content)
	}
	a.Messages = nil
}

func (a *RemoteActor) ReceiveMessage(m Message, reply *string) error {
	fmt.Printf("Remote actor %s received message: %s\n", m.To, m.Content)
	*reply = "message received"

	return nil
}