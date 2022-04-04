package captcha

import (
	"Zhigui/pkg/config"
	"Zhigui/pkg/redis"
	"Zhigui/pkg/web"
	"errors"
	"time"
)

// RedisStore 实现 base64Captcha.Store interface
type RedisStrore struct {
	RedisClinet *redis.RedisClient
	KeyPrefix   string
}

// Set 实现 base64Captcha.Store interface 的 Set 方法
func (s *RedisStrore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	// 方便本地开发调试
	if web.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.debug_expire_time"))
	}
	if ok := s.RedisClinet.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("无法存储图片验证码答案")
	}
	return nil
}

// Get 实现 base64Captcha.Store interface 的 Get 方法
func (s *RedisStrore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClinet.Get(key)
	if clear {
		s.RedisClinet.Del(key)
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStrore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
