package captcha

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

func NewDefaultRedisStore(redis *redis.Client, ctx context.Context) *RedisStore {
	return &RedisStore{
		Redis:      redis,
		Context:    ctx,
		Expiration: time.Second * 600,
		PreKey:     "CAPTCHA_",
	}
}

type RedisStore struct {
	Redis      *redis.Client
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (redisStore *RedisStore) Set(id string, value string) error {
	return redisStore.Redis.Set(redisStore.Context, redisStore.PreKey+id, value, redisStore.Expiration).Err()
}

func (redisStore *RedisStore) Get(key string, clear bool) string {
	val, err := redisStore.Redis.Get(redisStore.Context, key).Result()
	if err != nil {
		return ""
	}

	if clear {
		err := redisStore.Redis.Del(redisStore.Context, key).Err()
		if err != nil {
			return ""
		}
	}

	return val
}

func (redisStore *RedisStore) Verify(id, answer string, clear bool) bool {
	key := redisStore.PreKey + id
	v := redisStore.Get(key, clear)

	return v == answer
}
