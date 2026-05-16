package env

import (
	"github.com/caarlos0/env/v11"
)

type kafkaEnvConfig struct {
	Brokers     []string `env:"KAFKA_BROKERS,required"`
	EmailTopik  string   `env:"EMAIL_TOPIC,required"`
	EventsTopic string   `env:"EVENTS_TOPIC,required"`
	DLQTopic    string   `env:"DLQ_TOPIC,required"`
}

type kafkaConfig struct {
	kafkaEnvConfig
}

func NewKafkaConfig() (*kafkaConfig, error) {
	var kafkaEnvConfig kafkaEnvConfig
	err := env.Parse(&kafkaEnvConfig)
	if err != nil {
		return nil, err
	}
	return &kafkaConfig{kafkaEnvConfig: kafkaEnvConfig}, nil
}

func (cfg *kafkaConfig) GetBrokers() []string {
	return cfg.Brokers
}

func (cfg *kafkaConfig) GetEmailTopic() string {
	return cfg.EmailTopik
}

func (cfg *kafkaConfig) GetEventsTopic() string {
	return cfg.EventsTopic
}

func (cfg *kafkaConfig) GetDLQTopic() string {
	return cfg.DLQTopic
}
