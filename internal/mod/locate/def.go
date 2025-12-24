package locate

import (
	_ "embed"
)

const (
	Redis = "redis"
)

const (
	RedisPackage = "github.com/dawnsgo/dawn/locate/redis"
)

var (
	//go:embed redis.toml
	RedisTemplate string
)
