package bootstrap

import (
	"Zhigui/pkg/config"
	"Zhigui/pkg/redis"
	"fmt"
)

// SetupRedis 初始化Redis
func SetupRedis() {

	// 建立Redis连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"),
			config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
	)
}
