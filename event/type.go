package event

import (
	"context"
)

type Publisher interface {
	Publish(ctx context.Context, topic, key string, event Event) error
}

type Event interface {
	Name() string
}
