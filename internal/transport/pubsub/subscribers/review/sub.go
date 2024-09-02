package review

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"google.golang.org/api/option"
)

type Options struct {
	Filename         string
	ProjectID        string
	SubscriptionName string
}

func NewSubscriber(ctx context.Context, opt Options) (*Subscriber, error) {
	sa := option.WithCredentialsFile(opt.Filename)
	client, err := pubsub.NewClient(ctx, opt.ProjectID, sa)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pubsub client: %w", err)
	}

	sub := client.Subscription(opt.SubscriptionName)
	return &Subscriber{
		subscription: sub,
	}, nil
}

type Subscriber struct {
	subscription *pubsub.Subscription
}

func (s *Subscriber) Subscribe(ctx context.Context) error {
	err := s.subscription.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(string(m.Data))
		m.Ack()
	})
	if err != nil {
		return fmt.Errorf("failed to receive message from pub/sub: %w", err)
	}

	return nil
}
