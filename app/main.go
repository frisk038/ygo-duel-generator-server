package main

import (
	"context"
	"fmt"
	"ygo-generator-server/infra/repository"
)

func main() {
	//TODO
	// STUN SERVER
	// - GET ROOM_ID / GET body / POST Offer
	// - GET ROOM_ID / GET body / POST Offer ICE Candidate
	// - GET ROOM_ID / GET body / POST Answer ICE Candidate
	// - GET ROOM_ID / GET body / POST Answer
	//
	// GAMEPLAY
	// - GET LIFE POINTS
	// - SET LIFE POINTS

	repo, err := repository.NewClient(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	repo = repo
	for {

	}
}
