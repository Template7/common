package cache

import (
	"context"
	"github.com/Template7/common/config"
	"github.com/Template7/common/logger"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *redis.Client
)

func New() *redis.Client {
	once.Do(func() {
		instance = redis.NewClient(&redis.Options{
			Addr:         config.New().Cache.Host,
			Password:     config.New().Cache.Password,
			ReadTimeout:  time.Duration(config.New().Cache.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(config.New().Cache.WriteTimeout) * time.Second,
		})
		if err := instance.Ping(context.Background()).Err(); err != nil {
			logger.New().WithError(err).Panic("fail to ping redis")
			panic(err)
		}
		logger.New().Debug("common redis client initialized")
	})
	return instance
}
