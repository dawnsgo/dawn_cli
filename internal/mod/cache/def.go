package cache

import (
	_ "embed"
)

const (
	Redis    = "redis"
	Memcache = "memcache"
)

const (
	RedisPackage    = "github.com/dawnsgo/dawn/cache/redis"
	MemcachePackage = "github.com/dawnsgo/dawn/cache/memcache"
)

var (
	//go:embed redis.toml
	RedisTemplate string
	//go:embed memcache.toml
	MemcacheTemplate string
)
