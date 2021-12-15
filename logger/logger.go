package logger

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
	"strings"
	"sync"
)

var (
	once   sync.Once
	logger = logrus.New()

	// set log level
	logLevel = map[string]logrus.Level{
		"DEBUG": logrus.DebugLevel,
		"INFO":  logrus.InfoLevel,
		"WARN":  logrus.WarnLevel,
		"ERROR": logrus.ErrorLevel,
		"FATAL": logrus.FatalLevel,
		"PANIC": logrus.PanicLevel,
	}
)

func SetLevel(lvl string) {
	if l, exist := logLevel[lvl]; exist {
		logger.SetLevel(l)
	} else {
		logger.Warn("unsupported log level: ", lvl, ". use default default level: debug")
		logger.SetLevel(logrus.DebugLevel)
	}
}

func SetFormatter(fmter string) {
	timestampFormat := "2006-01-02 15:04:05.000"
	callerFormatter := func(path string) string {
		arr := strings.Split(path, "/")
		return arr[len(arr)-1]
	}
	callerPrettyfier := func(f *runtime.Frame) (string, string) {
		funcName := f.Func.Name()
		names := strings.Split(funcName, "/")
		return names[len(names)-1], fmt.Sprintf("%s:%d", callerFormatter(f.File), f.Line)
	}
	var formatter logrus.Formatter
	switch fmter {
	case "JSON":
		formatter = &logrus.JSONFormatter{
			TimestampFormat:  timestampFormat,
			CallerPrettyfier: callerPrettyfier,
		}
	case "STRING":
		formatter = &Formatter{
			TimestampFormat:  timestampFormat,
			CallerPrettyfier: callerPrettyfier,
		}
	default:
		logger.Warn("unsupported formatter: ", fmter, ". use default string formatter")
		formatter = &Formatter{
			TimestampFormat:  timestampFormat,
			CallerPrettyfier: callerPrettyfier,
		}
	}
	logger.SetFormatter(formatter)
}

func GetWithContext(ctx context.Context) *logrus.Entry {
	return logger.WithField("req_id", ctx.Value("X-Request-ID"))
}

func GetLogger() *logrus.Logger {
	once.Do(func() {
		logger.SetReportCaller(true)

		logger.Debug("logger initialized")
	})
	return logger
}
