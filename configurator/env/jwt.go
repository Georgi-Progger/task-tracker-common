package env

import (
	"github.com/caarlos0/env/v11"
)

type jwtEnvConfig struct {
	JWTSecret string `env:"JWT_SECRET,required"`
}

type jwtConfig struct {
	jwt jwtEnvConfig
}

func NewJWTConfig() (*jwtConfig, error) {
	var jwt jwtEnvConfig
	err := env.Parse(&jwt)
	if err != nil {
		return nil, err
	}
	return &jwtConfig{jwt: jwt}, nil
}

func (cfg *jwtConfig) GetJWTSecret() string {
	return cfg.jwt.JWTSecret
}
