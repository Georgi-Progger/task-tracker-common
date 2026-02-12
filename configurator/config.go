package configurator

import (
	"fmt"

	"github.com/Georgi-Progger/task-tracker-common/configurator/env"
	"github.com/joho/godotenv"
)

type Config struct {
	DbConfig
	AppConfig
	KafkaConfig
	RedisConfig
	JWTConfig
	EmailConfig
}

func LoadConfig() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("error loading .env file: %w", err)
	}

	db, err := env.NewDbConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error create db config: %w", err)
	}

	app, err := env.NewAppConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error create app config: %w", err)
	}

	kafka, err := env.NewKafkaConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error create kafka config: %w", err)
	}

	redis, err := env.NewRedisConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error create redis config: %w", err)
	}

	jwt, err := env.NewJWTConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error create jwt config: %w", err)
	}

	email, err := env.NewEmailConfig()
	if err != nil {
		return Config{}, fmt.Errorf("error create jwt config: %w", err)
	}

	cfg := Config{
		DbConfig:    db,
		AppConfig:   app,
		KafkaConfig: kafka,
		RedisConfig: redis,
		JWTConfig:   jwt,
		EmailConfig: email,
	}
	return cfg, nil
}
