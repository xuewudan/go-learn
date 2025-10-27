package config

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Config 应用配置
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort int
}

// Load 加载配置
func Load() *Config {
	return &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnvInt("DB_PORT", 3306),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "xuetao123"),
		DBName:     getEnv("DB_NAME", "blog_db"),
		JWTSecret:  getEnv("JWT_SECRET", "your_jwt_secret_key"),
		ServerPort: getEnvInt("SERVER_PORT", 8080),
	}
}

// 从环境变量获取字符串，默认值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 从环境变量获取整数，默认值
func getEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intVal, err := strconv.Atoi(value)
	if err != nil {
		logrus.Warnf("环境变量 %s 转换失败，使用默认值 %d", key, defaultValue)
		return defaultValue
	}
	return intVal
}
