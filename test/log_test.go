package test

import (
	"github.com/Template7/common/logger"
	"github.com/spf13/viper"
	"testing"
)

func init() {
	viper.AddConfigPath("./")
}

func TestLog(t *testing.T) {
	log := logger.New("debug", "console", "test")
	log.With("k1", "v1").Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	//log.Error("error msg")

	log = logger.GetLogger()
	log.Debug("debug msg")
}
