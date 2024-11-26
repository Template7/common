package t7Id

import (
	"github.com/Template7/common/logger"
	"github.com/bwmarrin/snowflake"
	"sync"
	"time"
)

var (
	once     sync.Once
	instance *snowflake.Node
)

func New() *snowflake.Node {
	once.Do(func() {
		node, err := snowflake.NewNode(time.Now().UnixNano() % 1024)
		if err != nil {
			panic(err)
		}
		instance = node

		logger.GetLogger().Info("t7 id initialized")
	})
	return instance
}
