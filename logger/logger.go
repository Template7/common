package logger

import (
	"fmt"
	"github.com/Template7/common/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"sync"
)

var (
	once     sync.Once
	instance *Logger
)

type Logger struct {
	core *zap.SugaredLogger
}

func (l *Logger) With(key string, value interface{}) *Logger {
	return &Logger{
		core: l.core.With(key, value),
	}
}

func (l *Logger) WithError(err error) *Logger {
	return &Logger{
		core: l.core.With("error", err.Error()),
	}
}

func (l *Logger) WithService(service string) *Logger {
	return &Logger{
		core: l.core.With("service", service),
	}
}

func (l *Logger) Debug(msg string) {
	l.core.Debug(msg)
}

func (l *Logger) Info(msg string) {
	l.core.Info(msg)
}

func (l *Logger) Warn(msg string) {
	l.core.Warn(msg)
}

func (l *Logger) Error(msg string) {
	l.core.Error(msg)
}

func (l *Logger) Panic(msg string) {
	l.core.Panic(msg)
}

func New() *Logger {
	once.Do(func() {
		cfg := config.New().Logger
		zCfg := zap.NewProductionConfig()
		zCfg.EncoderConfig.LevelKey = "log_level"
		zCfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

		lvlM := map[string]zapcore.Level{
			"debug": zap.DebugLevel,
			"info":  zap.InfoLevel,
			"warn":  zap.WarnLevel,
			"error": zap.ErrorLevel,
		}

		// set log level
		if lvl, exist := lvlM[cfg.Level]; exist {
			zCfg.Level = zap.NewAtomicLevelAt(lvl)
		} else {
			zCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
			log.Println(fmt.Sprintf("invalid log level: %s use debug level as default", cfg.Level))
		}

		logger, _ := zCfg.Build(zap.AddCallerSkip(1))
		defer logger.Sync() // flushes buffer, if any

		instance = &Logger{
			core: logger.Sugar().With("version", os.Getenv("GIT_TAG")),
		}

		instance.Info("logger initialized")
	})

	return instance
}
