package cache

import (
	"context"
	"github.com/Template7/common/logger"
	"github.com/redis/go-redis/v9"
	"time"
)

func New(host string, password string, readTimeout, writeTimeout int, log *logger.Logger) *redis.Client {
	log = log.WithContext(context.Background()).WithService("redisCore")

	if readTimeout == 0 {
		readTimeout = 3
	}
	if writeTimeout == 0 {
		writeTimeout = 3
	}

	instance := redis.NewClient(&redis.Options{
		Addr:         host,
		Password:     password,
		ReadTimeout:  time.Duration(readTimeout) * time.Second,
		WriteTimeout: time.Duration(writeTimeout) * time.Second,
	})
	if err := instance.Ping(context.Background()).Err(); err != nil {
		log.WithError(err).Panic("fail to ping redis")
		panic(err)
	}
	log.Info("common redis client initialized")
	return instance
}
