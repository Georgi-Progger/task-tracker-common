package configurator

type DbConfig interface {
	GetUrlDb() string
}

type RedisConfig interface {
	GetUrlRedis() string
}

type AppConfig interface {
	GetPort() string
}

type BrokerConfig interface {
	GetBrokers() []string
}

type JWTConfig interface {
	GetJWTSecret() string
}
