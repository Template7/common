package db

import (
	"context"
	"fmt"
	"github.com/Template7/common/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewNoSql(host string, port int, username string, password string) *mongo.Client {
	var cs string
	if username != "" && password != "" {
		cs = fmt.Sprintf("mongodb://%s:%s@%s:%d", username, password, host, port)
	} else {
		cs = fmt.Sprintf("mongodb://%s:%d", host, port)
	}
	c, err := mongo.Connect(nil, options.Client().ApplyURI(cs))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := c.Ping(ctx, nil); err != nil {
		panic(err)
	}

	logger.GetLogger().Info("mongo client initialized")
	return c
}
