package types

import (
	"fmt"
	"net/rpc"
)

type IActorSystem interface {
	StartSystem() *ActorSystem
	CreateActor(name string)
	SendMessage(m Message)
}

type ActorSystem struct {
	Actors map[string]*Actor
}

func StartSystem() *ActorSystem {
	return &ActorSystem{
		Actors: make(map[string]*Actor),
	}
}

func (s *ActorSystem) CreateActor(name string) {
	s.Actors[name] = &Actor{Name: name}
}

func (s *ActorSystem) SendMessage(m Message) {
	
	if actor, ok := s.Actors[m.To]; ok {
		actor.SendMessage(m)
	} else {
		client, err := rpc.Dial("tcp", m.To)

		if err != nil {
			fmt.Printf("Error connecting to remote actor %s: %s\n", m.To, err)
			return
		}

		defer client.Close()

		var reply string

		err = client.Call("RemoteActor.ReceiveMessage", m, &reply)

		if err != nil {
			fmt.Printf("error sending message to remote actor %s: %s\n", m.To, err)
			return
		}
	}
}