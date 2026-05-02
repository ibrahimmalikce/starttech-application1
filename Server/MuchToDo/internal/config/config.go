package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	ServerPort         string `mapstructure:"port"`
	MongoURI           string `mapstructure:"mongo_uri"`
	DBName             string `mapstructure:"db_name"`
	JWTSecretKey       string `mapstructure:"jwt_secret_key"`
	JWTExpirationHours int    `mapstructure:"jwt_expiration_hours"`
	EnableCache        bool   `mapstructure:"enable_cache"`
	RedisAddr          string `mapstructure:"redis_addr"`
	RedisPassword      string `mapstructure:"redis_password"`
	LogLevel           string `mapstructure:"log_level"`
	LogFormat          string `mapstructure:"log_format"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.SetDefault("port", getEnv("PORT", "8080"))
	viper.SetDefault("mongo_uri", getEnv("MONGO_URI", ""))
	viper.SetDefault("db_name", getEnv("DB_NAME", "much_todo_db"))
	viper.SetDefault("jwt_secret_key", getEnv("JWT_SECRET_KEY", ""))
	viper.SetDefault("jwt_expiration_hours", 72)
	viper.SetDefault("enable_cache", false)
	viper.SetDefault("log_level", getEnv("LOG_LEVEL", "DEBUG"))
	viper.SetDefault("log_format", getEnv("LOG_FORMAT", "json"))

	err = viper.Unmarshal(&config)
	return
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
