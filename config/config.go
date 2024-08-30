package config

import (
	"os"
)

// GetEnv 获取环境变量值，若环境变量未设置则返回默认值
func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

var (
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func LoadConfig() {
	DBHost = GetEnv("DB_HOST", "localhost") // "localhost"
	DBPort = GetEnv("DB_PORT", "5432")
	DBUser = GetEnv("DB_USER", "tarot")
	DBPassword = GetEnv("DB_PASSWORD", "tarot_password")
	DBName = GetEnv("DB_NAME", "tarot")
}
