package config

import (
	"github.com/spf13/viper"
	"log"
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
	Db struct {
		Sql struct {
			Db         string
			Host       string
			Port       int
			Username   string
			Password   string
			Connection struct {
				Min int
				Max int
			}
		}
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
			panic(err)
		}
		if err := viper.Unmarshal(&instance); err != nil {
			panic(err)
		}

		log.Println("common config initialized")
	})
	return instance
}
