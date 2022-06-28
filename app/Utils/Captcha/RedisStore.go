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

func (this *RedisStore) Set(id string, value string) error {
	return this.Redis.Set(this.Context, this.PreKey+id, value, this.Expiration).Err()
}

func (this *RedisStore) Get(key string, clear bool) string {
	val, err := this.Redis.Get(this.Context, key).Result()
	if err != nil {
		return ""
	}

	if clear {
		err := this.Redis.Del(this.Context, key).Err()
		if err != nil {
			return ""
		}
	}
	return val
}

func (this *RedisStore) Verify(id, answer string, clear bool) bool {
	key := this.PreKey + id
	v := this.Get(key, clear)
	return v == answer
}
