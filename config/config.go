package config

import (
	"github.com/Template7/common/logger"
	"github.com/spf13/viper"
	"sync"
)

const (
	configPath = "config"
)

type Config struct {
	Logger struct {
		Level string
	}
	Cache struct {
		Host         string
		Password     string
		ReadTimeout  int
		WriteTimeout int
	}
}

var (
	once     sync.Once
	instance *Config
)

func New() *Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		instance = &Config{}
		viper.AddConfigPath(configPath)
		viper.SetConfigName("config")
		if err := viper.ReadInConfig(); err != nil {
			logger.New().WithError(err).Panic("fail to read config")
			panic(err)
		}
		if err := viper.Unmarshal(&instance); err != nil {
			logger.New().WithError(err).Panic("fail to unmarshal config")
			panic(err)
		}

		logger.New().Info("common config initialized")
	})
	return instance
}
