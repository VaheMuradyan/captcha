package captchaimage

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisStore struct {
	client  *redis.Client
	timeout time.Duration
	ctx     context.Context
}

func NewRedisStore(client *redis.Client, timeout time.Duration, ctx context.Context) *RedisStore {
	return &RedisStore{
		client:  client,
		timeout: timeout,
		ctx:     ctx,
	}
}

func (s *RedisStore) Set(id string, digiths []byte) {

	if id == "" || len(digiths) == 0 {
		log.Println("invalid CAPTCHA data: ID or digiths are empty")
		return
	}

	captchaSolution := ""
	for _, digit := range digiths {
		captchaSolution += strconv.Itoa(int(digit))
	}

	err := s.client.Set(s.ctx, id, captchaSolution, s.timeout).Err()
	if err != nil {
		panic("Faild to set CAPTCHA in Redis " + err.Error())
	}

}
func (s *RedisStore) Get(id string, clear bool) []byte {
	result, err := s.client.Get(s.ctx, id).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		panic("Failed to get CAPTCHA from Redis " + err.Error())
	}

	if clear {
		_, _ = s.client.Del(s.ctx, id).Result()
	}

	digits := []byte{}
	for _, char := range result {
		digits = append(digits, byte(char-'0'))
	}
	return digits
}
