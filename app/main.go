package main

import (
	"context"
	"fmt"
	v1 "ygo-generator-server/app/handler/v1"
	"ygo-generator-server/business"
	"ygo-generator-server/infra/repository"

	"github.com/gin-gonic/gin"
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

	ctx := context.Background()
	repo, err := repository.NewClient(ctx)
	if err != nil {
		fmt.Println(err)
	}
	stunb := business.NewStunBusiness(repo)
	stunh := v1.NewStunHandler(stunb)

	r := gin.Default()
	stunRoute := r.Group("/stun")
	stunRoute.GET("/offer/:roomid", stunh.GetOfferHandler)
	stunRoute.GET("/offerice/:roomid", stunh.GetOfferICEHandler)
	stunRoute.GET("/answer/:roomid", stunh.GetAnswerHandler)
	stunRoute.GET("/answerice/:roomid", stunh.GetAnswerICEHandler)
	stunRoute.POST("/offer/:roomid", stunh.PostOfferHandler)
	stunRoute.POST("/offerice/:roomid", stunh.PostOfferICEHandler)
	stunRoute.POST("/answer/:roomid", stunh.PostAnswerHandler)
	stunRoute.POST("/answerice/:roomid", stunh.PostAnswerICEHandler)

	r.Run()
}
