package configurator

type DbConfig interface {
	GetUrlDb() string
}

type RedisConfig interface {
	GetUrlRedis() string
}

type AppConfig interface {
	GetServerPort() string
}

type ShedulerConfig interface {
	GetSchedulerHost() string
	GetSchedulerPort() string
}

type KafkaConfig interface {
	GetBrokers() []string
	GetEmailTopic() string
	GetEventsTopic() string
	GetDLQTopic() string
}

type JWTConfig interface {
	GetJWTSecret() string
}

type EmailConfig interface {
	GetEmailHost() string
	GetEmailLogin() string
	GetEmailPassword() string
	GetEmailPort() int
}
