package env

import (
	"github.com/caarlos0/env/v11"
)

type brokerEnvConfig struct {
	Brokers []string `env:"KAFKA_BROKERS,required"`
}

type brokerConfig struct {
	broker brokerEnvConfig
}

func NewBrokerConfig() (*brokerConfig, error) {
	var broker brokerEnvConfig
	err := env.Parse(&broker)
	if err != nil {
		return nil, err
	}
	return &brokerConfig{broker: broker}, nil
}

func (cfg *brokerConfig) GetBrokers() []string {
	return cfg.broker.Brokers
}
