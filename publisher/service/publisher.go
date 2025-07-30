package publisher

import (
	"context"

	"github.com/aldadelas/pub-sub-go/client"
)

type Publisher struct {
	redisClient *client.RedisClient
}

func NewPublisher(redisClient *client.RedisClient) *Publisher {
	return &Publisher{redisClient: redisClient}
}

func (p *Publisher) Publish(ctx context.Context, channel string, message string) error {
	return p.redisClient.Client.Publish(ctx, channel, message).Err()
}
