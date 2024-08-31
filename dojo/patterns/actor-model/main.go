package main

import "github.com/Alastair7/knowledge-journey/actor-model/types"

func main() {
	actorSys := types.StartSystem()
	
	actorSys.CreateActor("actor1")
	actorSys.CreateActor("actor2")

	actorSys.SendMessage(types.Message{
		To: "actor2",
		From: "actor1",
		Content: "Hey! This is actor1",
	})

	actorSys.SendMessage(types.Message{
		To: "actor1",
		From: "actor2",
		Content: "Hey! This is actor2",
	})

	for _, actor := range actorSys.Actors {
		actor.ProcessMessages()
	}
}