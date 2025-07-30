package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aldadelas/pub-sub-go/client"
	publisher "github.com/aldadelas/pub-sub-go/publisher/service"
)

func main() {
	ctx := context.Background()
	redisClient := client.NewRedisClient("localhost:6379", "", 0)
	publisher := publisher.NewPublisher(redisClient)

	for i := 0; i < 3; i++ {
		publisher.Publish(ctx, "test", fmt.Sprintf("Hello, World! %d", i+1))
		time.Sleep(1 * time.Second)
	}
	redisClient.Client.Close()
}
