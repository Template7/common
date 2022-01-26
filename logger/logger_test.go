package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"testing"
)

func TestGetLogger(t *testing.T) {
	viper.AddConfigPath("../test")
	log := GetLogger()
	SetLevel("DEBUG")
	SetFormatter("STRING")
	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")

	SetLevel("INFO")
	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")

	SetLevel("WARN")
	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")

	SetFormatter("JSON")
	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")

	SetLevel("INFO")
	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")

	SetLevel("WARN")
	log.Debug("debug log")
	log.Info("info log")
	log.Warn("warn log")
	log.Error("error log")
}

// go test -bench ^BenchmarkLogging -run none -benchmem ./logger
func BenchmarkLogging(b *testing.B) {
	viper.AddConfigPath("../test")
	SetLevel("DEBUG")
	SetFormatter("STRING")
	log := GetLogger()
	//time.Sleep(10 * time.Second)
	for i := 0; i < b.N; i++ {
		log.Debug("debug log")
	}
}

func BenchmarkLoggingWithContext(b *testing.B) {
	viper.AddConfigPath("../test")
	SetLevel("DEBUG")
	SetFormatter("STRING")
	ctx := gin.Context{}
	ctx.Set("X-Request-ID", "requestId")

	for i := 0; i < b.N; i++ {
		log := GetWithContext(&ctx)
		log.Debug("debug log")
	}
}
