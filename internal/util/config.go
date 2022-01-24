package util

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	ServerAddress              string
	DBHost                     string
	DBName                     string
	DBUser                     string
	DBPass                     string
	DBPort                     string
	AccessTokenPrivateKeyPath  string
	AccessTokenPublicKeyPath   string
	RefreshTokenPrivateKeyPath string
	RefreshTokenPublicKeyPath  string
	JwtExpiration              int // in minutes
}

func NewConfiguration() *Configuration {
	viper.AutomaticEnv()

	viper.SetDefault("SERVER_ADDRESS", "0.0.0.0:8080")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_NAME", "basic-api")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_PASSWORD", "postgres")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("ACCESS_TOKEN_PRIVATE_KEY_PATH", "./certs/access-private.pem")
	viper.SetDefault("ACCESS_TOKEN_PUBLIC_KEY_PATH", "./certs/access-public.pem")
	viper.SetDefault("REFRESH_TOKEN_PRIVATE_KEY_PATH", "./certs/refresh-private.pem")
	viper.SetDefault("REFRESH_TOKEN_PUBLIC_KEY_PATH", "./certs/refresh-public.pem")
	viper.SetDefault("JWT_EXPIRATION", 30)

	config := &Configuration{
		ServerAddress:              viper.GetString("SERVER_ADDRESS"),
		DBHost:                     viper.GetString("DB_HOST"),
		DBName:                     viper.GetString("DB_NAME"),
		DBUser:                     viper.GetString("DB_USER"),
		DBPass:                     viper.GetString("DB_PASSWORD"),
		DBPort:                     viper.GetString("DB_PORT"),
		JwtExpiration:              viper.GetInt("JWT_EXPIRATION"),
		AccessTokenPrivateKeyPath:  viper.GetString("ACCESS_TOKEN_PRIVATE_KEY_PATH"),
		AccessTokenPublicKeyPath:   viper.GetString("ACCESS_TOKEN_PUBLIC_KEY_PATH"),
		RefreshTokenPrivateKeyPath: viper.GetString("REFRESH_TOKEN_PRIVATE_KEY_PATH"),
		RefreshTokenPublicKeyPath:  viper.GetString("REFRESH_TOKEN_PUBLIC_KEY_PATH"),
	}

	return config
}
