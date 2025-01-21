package logger

import (
	"os"

	"go.elastic.co/ecszap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type BuilderOption struct {
	PrettyPrint bool
}

func getLogLevel(logLevel string) zap.AtomicLevel {
	switch logLevel {
	case "debug":
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case "info":
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case "warn":
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case "error":
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	}
}

func NewLogger(serviceName string, level string, options ...func(*BuilderOption)) *zap.Logger {
	// build config
	cfg := &BuilderOption{}
	for _, option := range options {
		option(cfg)
	}

	// create new core with log duplication
	var core zapcore.Core
	if cfg.PrettyPrint {
		// create new console formatter config with colored level
		config := zap.NewDevelopmentEncoderConfig()
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
		core = zapcore.NewCore(zapcore.NewConsoleEncoder(config), zapcore.AddSync(os.Stdout), getLogLevel(level))
	} else {
		// create new ECS formatter config
		core = ecszap.NewCore(ecszap.NewDefaultEncoderConfig(), zapcore.AddSync(os.Stdout), getLogLevel(level))
	}

	// create new log instance
	log := zap.New(core, zap.AddCaller())
	log = log.With(zap.String("service_name", serviceName))

	return log
}

func WithID(log *zap.Logger, contextName string, scopeName string) *zap.Logger {
	return log.With(zap.String("context", contextName), zap.String("scope", scopeName))
}

func WithPrettyPrint() func(*BuilderOption) {
	return func(o *BuilderOption) {
		o.PrettyPrint = true
	}
}
