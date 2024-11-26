package db

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"

	"github.com/Template7/common/logger"
	"gorm.io/gorm"
)

func NewSql(host string, port int, username string, password string, db string, idleConn, openConn int, log *logger.Logger) *gorm.DB {
	log = log.WithContext(context.Background()).WithService("sqlCore")

	cs := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC", username, password, host, port, db)
	sqlDb, err := gorm.Open(mysql.Open(cs), &gorm.Config{})
	if err != nil {
		log.WithError(err).With("connection_string", cs).Panic("fail to connect to db")
		panic(err)
	}
	conn, err := sqlDb.DB()
	if err != nil {
		log.WithError(err).Panic("fail to get db connection")
		panic(err)
	}
	if err := conn.Ping(); err != nil {
		log.WithError(err).Panic("fail to ping db")
		panic(err)
	}

	if idleConn == 0 {
		idleConn = 4
	}
	if openConn == 0 {
		openConn = 8
	}
	conn.SetMaxIdleConns(idleConn)
	conn.SetMaxOpenConns(openConn)

	log.Info("common sql initialized")
	return sqlDb
}
