package review

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type PubOptions struct {
	Filename  string
	ProjectID string
	Topic     string
}

func NewPublisher(ctx context.Context, opt PubOptions) (*Publisher, error) {
	sa := option.WithCredentialsFile(opt.Filename)
	client, err := pubsub.NewClient(ctx, opt.ProjectID, sa)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pub/sub client: %w", err)
	}

	topic := client.Topic(opt.Topic)
	return &Publisher{
		topic: topic,
	}, nil
}

type Publisher struct {
	topic *pubsub.Topic
}

type UserAddedMessage struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (p *Publisher) Publish(ctx context.Context, msg UserAddedMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshall user added msg: %w", err)
	}

	_ = p.topic.Publish(ctx, &pubsub.Message{
		Data: data,
	})
	return nil
}
