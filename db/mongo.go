package db

import (
	"context"
	"fmt"
	"github.com/Template7/common/config"
	"github.com/Template7/common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

var (
	mOnce     sync.Once
	mInstance *mongo.Client
)

func NewNoSql() *mongo.Client {
	mOnce.Do(func() {
		var cs string
		cfg := config.New()
		if cfg.Db.NoSql.Username != "" && cfg.Db.NoSql.Password != "" {
			cs = fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.Db.NoSql.Username, cfg.Db.NoSql.Password, cfg.Db.NoSql.Host, cfg.Db.NoSql.Port)
		} else {
			cs = fmt.Sprintf("mongodb://%s:%d", cfg.Db.NoSql.Host, cfg.Db.NoSql.Port)
		}
		c, err := mongo.Connect(nil, options.Client().ApplyURI(cs))
		if err != nil {
			panic(err)
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		if err := c.Ping(ctx, nil); err != nil {
			cancel()
			panic(err)
		}
		mInstance = c

		logger.New().Info("mongo client initialized")
	})
	return mInstance
}
