package cache

import (
   "gopkg.in/redis.v3"
   "os"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache()*RedisCache {
	rc := new(RedisCache)
	rc.client = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR")
	})
	return rc
}

func (rc *RedisCache) Put(key, value string) error{
	return nil
}

func(rc *RedisCache) Get(key) (string, error) {
	return ""
}