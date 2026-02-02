package kafka

import "context"

type Handler func(ctx context.Context, msg []byte) error

type Producer interface {
	Send(ctx context.Context, value []byte) error
	Close() error
}

type Consumer interface {
	Start(ctx context.Context, handler Handler) error
	Close() error
}
