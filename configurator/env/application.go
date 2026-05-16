package env

import "github.com/caarlos0/env/v11"

type appEnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
}
type appConfig struct {
	appEnvConfig
}

func NewAppConfig() (*appConfig, error) {
	var cfg appEnvConfig
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &appConfig{appEnvConfig: cfg}, nil
}

func (cfg *appConfig) GetServerPort() string {
	return cfg.ServerPort
}
