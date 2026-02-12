package env

import (
	"github.com/caarlos0/env/v11"
)

type kafkaEnvConfig struct {
	Brokers     []string `env:"KAFKA_BROKERS,required"`
	EmailTopik  string   `env:"EMAIL_TOPIC,required"`
	EventsTopic string   `env:"EVENTS_TOPIC,required"`
}

type kafkaConfig struct {
	broker kafkaEnvConfig
}

func NewKafkaConfig() (*kafkaConfig, error) {
	var broker kafkaEnvConfig
	err := env.Parse(&broker)
	if err != nil {
		return nil, err
	}
	return &kafkaConfig{broker: broker}, nil
}

func (cfg *kafkaConfig) GetBrokers() []string {
	return cfg.broker.Brokers
}

func (cfg *kafkaConfig) GetEmailTopic() string {
	return cfg.broker.EmailTopik
}

func (cfg *kafkaConfig) GetEventsTopic() string {
	return cfg.broker.EventsTopic
}
