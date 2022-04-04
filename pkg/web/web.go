package web

import (
	"Zhigui/pkg/config"
)

func IsLocal() bool {
	return config.Get("web.env") == "local"
}

func IsProduction() bool {
	return config.Get("web.env") == "production"
}

func IsTesting() bool {
	return config.Get("web.env") == "testing"
}
