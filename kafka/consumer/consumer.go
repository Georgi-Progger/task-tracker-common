package consumer

import (
	"context"
	"time"

	kafkain "github.com/Georgi-Progger/task-tracker-common/kafka"
	"github.com/Georgi-Progger/task-tracker-common/logger"

	"github.com/segmentio/kafka-go"
)

type consumer struct {
	reader *kafka.Reader
	logger logger.Logger
}

func NewConsumer(dsn []string, topic string, logger logger.Logger) kafkain.Consumer {
	return &consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers: dsn,
			Topic:   topic,
			GroupID: "email-senders",
		}),
		logger: logger,
	}
}

func (c *consumer) Start(ctx context.Context, handler kafkain.Handler) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				c.logger.Error(err, "read error")
				time.Sleep(time.Second)
				continue
			}

			if err := handler(ctx, msg.Value); err != nil {
				c.logger.Error(err, "handler failed")
			}
		}
	}
}

func (c *consumer) Close() error {
	return c.reader.Close()
}
