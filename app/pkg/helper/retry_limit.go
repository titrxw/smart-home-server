package helper

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const ServiceRetryLimit = "service:retry:limit:%s"

type RetryTooManyError struct {
	error
	Msg string
}

func (retryTooManyError RetryTooManyError) Error() string {
	return retryTooManyError.Msg
}

type RetryLimit struct {
}

func (retryLimit RetryLimit) Try(ctx context.Context, redisClient redis.Cmdable, callback func(context.Context, redis.Cmdable, int64) error, key string, tryNum int64, ttl time.Duration) error {
	curNum, err := retryLimit.incr(ctx, redisClient, key, ttl)
	if err != nil {
		return err
	}
	if curNum > tryNum {
		_ = retryLimit.Reset(ctx, redisClient, key)
		redisClient.Del(ctx, key) //极端情况下，会导致尝试次数内得请求获取不到redis数据

		return &RetryTooManyError{Msg: "失败次数过多，请稍后刷新重试"}
	}

	return callback(ctx, redisClient, curNum)
}

func (retryLimit RetryLimit) Reset(ctx context.Context, redisClient redis.Cmdable, key string) error {
	return redisClient.Del(ctx, retryLimit.formatCacheKey(key)).Err()
}

func (retryLimit RetryLimit) Decr(ctx context.Context, redisClient redis.Cmdable, key string) error {
	return redisClient.Decr(ctx, retryLimit.formatCacheKey(key)).Err()
}

func (retryLimit RetryLimit) formatCacheKey(key string) string {
	return fmt.Sprintf(ServiceRetryLimit, key)
}

func (retryLimit RetryLimit) incr(ctx context.Context, redisClient redis.Cmdable, key string, ttl time.Duration) (int64, error) {
	cacheKey := retryLimit.formatCacheKey(key)
	_, err := redisClient.SetNX(ctx, cacheKey, 0, ttl).Result()
	if err != nil {
		return -1, err
	}

	return redisClient.Incr(ctx, cacheKey).Result()
}
