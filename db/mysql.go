package db

import (
	"fmt"
	"github.com/Template7/common/config"
	"github.com/Template7/common/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var (
	once     sync.Once
	instance *gorm.DB
)

func NewSql() *gorm.DB {
	once.Do(func() {
		cfg := config.New()
		cs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Db.Sql.Username, cfg.Db.Sql.Password, cfg.Db.Sql.Host, cfg.Db.Sql.Port, cfg.Db.Sql.Db)
		sqlDb, err := gorm.Open(mysql.Open(cs), &gorm.Config{})
		if err != nil {
			logger.New().WithError(err).With("connection_string", cs).Panic("fail to connect to db")
			panic(err)
		}
		conn, err := sqlDb.DB()
		if err != nil {
			logger.New().WithError(err).Panic("fail to get db connection")
			panic(err)
		}
		if err := conn.Ping(); err != nil {
			logger.New().WithError(err).Panic("fail to ping db")
			panic(err)
		}

		conn.SetMaxIdleConns(cfg.Db.Sql.Connection.Min)
		conn.SetMaxOpenConns(cfg.Db.Sql.Connection.Max)
		instance = sqlDb
	})

	return instance
}
