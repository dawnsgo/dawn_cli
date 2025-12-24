package lock

import (
	_ "embed"
)

const (
	Redis    = "redis"
	Memcache = "memcache"
)

const (
	RedisPackage    = "github.com/dawnsgo/dawn/lock/redis"
	MemcachePackage = "github.com/dawnsgo/dawn/lock/memcache"
)

var (
	//go:embed redis.toml
	RedisTemplate string
	//go:embed memcache.toml
	MemcacheTemplate string
)
