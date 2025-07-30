package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aldadelas/pub-sub-go/client"
	subscriber "github.com/aldadelas/pub-sub-go/subscriber/service"
)

func main() {
	redisClient := client.NewRedisClient("localhost:6379", "", 0)
	subscriber := subscriber.NewSubscriber(redisClient)

	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(10*time.Second, func() {
		cancel()
	})

	err := subscriber.Subscribe(ctx, "test")
	if err != nil {
		fmt.Println("Error subscribing:", err)
	}

	redisClient.Client.Close()
}
