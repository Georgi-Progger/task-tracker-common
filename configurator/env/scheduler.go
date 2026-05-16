package env

import "github.com/caarlos0/env/v11"

type schedulerEnvConfig struct {
	SchedulerPort string `env:"SCHEDULER_PORT,required"`
	SchedulerHost string `env:"SCHEDULER_HOST,required"`
}

type schedulerConfig struct {
	schedulerEnvConfig
}

func NewSchedulerConfig() (*schedulerConfig, error) {
	var scheduler schedulerEnvConfig
	err := env.Parse(&scheduler)
	if err != nil {
		return nil, err
	}
	return &schedulerConfig{schedulerEnvConfig: scheduler}, nil
}

func (cfg *schedulerConfig) GetSchedulerHost() string {
	return cfg.SchedulerHost
}

func (cfg *schedulerConfig) GetSchedulerPort() string {
	return cfg.SchedulerPort
}
