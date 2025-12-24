package eventbus

import (
	_ "embed"
)

const (
	Nats    = "nats"
	Redis   = "redis"
	Kafka   = "kafka"
	Process = "process"
)

const (
	RedisPackage = "github.com/dawnsgo/dawn/eventbus/redis"
	NatsPackage  = "github.com/dawnsgo/dawn/eventbus/nats"
	KafkaPackage = "github.com/dawnsgo/dawn/eventbus/kafka"
)

var (
	//go:embed nats.toml
	NatsTemplate string
	//go:embed redis.toml
	RedisTemplate string
	//go:embed kafka.toml
	KafkaTemplate string
)
