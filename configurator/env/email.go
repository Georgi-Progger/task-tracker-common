package env

import "github.com/caarlos0/env/v11"

type emailEnvConfig struct {
	Host     string `env:"EMAIL_HOST,required"`
	Login    string `env:"EMAIL_LOGIN,required"`
	Password string `env:"EMAIL_PASSWORD,required"`
	Port     int    `env:"EMAIL_PORT,required"`
}

type emailConfig struct {
	email emailEnvConfig
}

func NewEmailConfig() (*emailConfig, error) {
	var email emailEnvConfig
	err := env.Parse(&email)
	if err != nil {
		return nil, err
	}
	return &emailConfig{email: email}, nil
}

func (cfg *emailConfig) GetEmailHost() string {
	return cfg.email.Host
}

func (cfg *emailConfig) GetEmailLogin() string {
	return cfg.email.Login
}

func (cfg *emailConfig) GetEmailPassword() string {
	return cfg.email.Password
}

func (cfg *emailConfig) GetEmailPort() int {
	return cfg.email.Port
}
