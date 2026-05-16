package producer

import (
	"context"
	"fmt"

	kafkain "github.com/Georgi-Progger/task-tracker-common/kafka"
	"github.com/Georgi-Progger/task-tracker-common/logger"

	"github.com/segmentio/kafka-go"
)

type producer struct {
	writer *kafka.Writer
	logger logger.Logger
}

func NewProducer(dsn []string, topic string, logger logger.Logger) kafkain.Producer {
	return &producer{
		writer: kafka.NewWriter(kafka.WriterConfig{
			Brokers: dsn,
			Topic:   topic,
		}),
		logger: logger,
	}
}

func (p *producer) Send(ctx context.Context, value []byte) error {

	msg := kafka.Message{
		Value: []byte(value),
	}

	err := p.writer.WriteMessages(ctx, msg)
	if err != nil {
		p.logger.Error(err, "could not write message")
		return err
	}

	p.logger.Info(fmt.Sprintf("Produced message: %s", value))
	return nil
}

func (p *producer) Close() error {
	return p.writer.Close()
}
