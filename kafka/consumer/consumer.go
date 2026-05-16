package consumer

import (
	"context"
	"time"

	kafkain "github.com/Georgi-Progger/task-tracker-common/kafka"
	"github.com/Georgi-Progger/task-tracker-common/logger"

	"github.com/segmentio/kafka-go"
)

type consumer struct {
	reader     *kafka.Reader
	dlqWriter  *kafka.Writer
	maxRetries uint
	logger     logger.Logger
}

func NewConsumer(dsn []string, topic, dlqTopic string, logger logger.Logger) kafkain.Consumer {
	return &consumer{
		reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:        dsn,
			Topic:          topic,
			GroupID:        "email-senders",
			CommitInterval: 100 * time.Millisecond,
		}),
		dlqWriter: kafka.NewWriter(kafka.WriterConfig{
			Brokers: dsn,
			Topic:   dlqTopic,
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
		}

		msg, err := c.reader.FetchMessage(ctx)
		if err != nil {
			c.logger.Error(err, "read error")
			continue
		}

		if err := handler(ctx, msg.Value); err != nil {
			c.logger.Error(err, "handler failed")
			if err = c.dlqWriter.WriteMessages(ctx, msg); err != nil {
				c.logger.Error(err, "failed to send to DLQ")
				continue
			}

			if err := c.reader.CommitMessages(ctx, msg); err != nil {
				c.logger.Error(err, "commit failed")
			}
			continue
		}

		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			c.logger.Error(err, "commit failed")
		}
	}
}

func (c *consumer) Close() error {
	return c.reader.Close()
}
