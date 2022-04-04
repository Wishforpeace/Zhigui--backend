// Package config 站点配置信息
package config

import (
	"Zhigui/pkg/config"
)

func init() {
	config.Add("app", func() map[string]interface{} {
		return map[string]interface{}{

			// 应用名称
			"name": config.Env("Web_NAME", "Zhigui"),

			// 当前环境，用以区分多环境，一般为 local, stage, production, test
			"env": config.Env("Web_ENV", "production"),

			// 是否进入调试模式
			"debug": config.Env("Web_DEBUG", false),

			// 应用服务端口
			"port": config.Env("Web_PORT", "8848"),

			// 加密会话、JWT 加密
			"key": config.Env("Web_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

			// 用以生成链接
			"url": config.Env("Web_URL", "http://localhost:3000"),

			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": config.Env("TIMEZONE", "Asia/Wuhan"),
		}
	})
}
