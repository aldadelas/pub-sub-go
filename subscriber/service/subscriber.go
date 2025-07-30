package subscriber

import (
	"context"
	"fmt"

	"github.com/aldadelas/pub-sub-go/client"
)

type Subscriber struct {
	redisClient *client.RedisClient
}

func NewSubscriber(redisClient *client.RedisClient) *Subscriber {
	return &Subscriber{redisClient: redisClient}
}

func (s *Subscriber) Subscribe(ctx context.Context, channel string) error {
	pubsub := s.redisClient.Client.Subscribe(ctx, channel)
	defer pubsub.Close()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case msg := <-pubsub.Channel():
			fmt.Println(msg.Payload)
		}
	}
}
