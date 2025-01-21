package database

import (
	"context"
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type (
	ContextFn func(ctx context.Context) []zapcore.Field

	CustomGormLogger struct {
		LogSQL                    bool
		ZapLogger                 *zap.Logger
		LogLevel                  gormlogger.LogLevel
		SlowThreshold             time.Duration
		SkipCallerLookup          bool
		IgnoreRecordNotFoundError bool
		Context                   ContextFn
	}
)

var (
	gormPackage    = filepath.Join("gorm.io", "gorm")
	zapgormPackage = filepath.Join("moul.io", "zapgorm2")
)

func NewGormLogger(zapLogger *zap.Logger, level gormlogger.LogLevel, logSql bool) CustomGormLogger {
	return CustomGormLogger{
		LogSQL:                    logSql,
		ZapLogger:                 zapLogger,
		LogLevel:                  level,
		SlowThreshold:             100 * time.Millisecond,
		SkipCallerLookup:          false,
		IgnoreRecordNotFoundError: false,
		Context:                   nil,
	}
}

func (l CustomGormLogger) SetAsDefault() {
	gormlogger.Default = l
}

func (l CustomGormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return CustomGormLogger{
		ZapLogger:                 l.ZapLogger,
		SlowThreshold:             l.SlowThreshold,
		LogLevel:                  level,
		SkipCallerLookup:          l.SkipCallerLookup,
		IgnoreRecordNotFoundError: l.IgnoreRecordNotFoundError,
		Context:                   l.Context,
	}
}

func (l CustomGormLogger) Info(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Info {
		return
	}
	l.logger(ctx).Sugar().Debugf(str, args...)
}

func (l CustomGormLogger) Warn(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Warn {
		return
	}
	l.logger(ctx).Sugar().Warnf(str, args...)
}

func (l CustomGormLogger) Error(ctx context.Context, str string, args ...interface{}) {
	if l.LogLevel < gormlogger.Error {
		return
	}
	l.logger(ctx).Sugar().Errorf(str, args...)
}

func (l CustomGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// if the log level is below 0, return
	if l.LogLevel <= 0 {
		return
	}

	// get query metadata
	elapsed := time.Since(begin)
	logger := l.logger(ctx)
	sql, rows := fc()

	// create log options
	logOpts := []zapcore.Field{
		zap.Duration("elapsed", elapsed),
		zap.Int64("rows", rows),
	}

	// check if we want to log the full SQL
	if l.LogSQL {
		logOpts = append(logOpts, zap.String("sql", sql))
	}

	// check if we want to log the error
	if err != nil && l.LogLevel >= gormlogger.Error && (!l.IgnoreRecordNotFoundError || !errors.Is(err, gorm.ErrRecordNotFound)) {
		logger.Error("trace", append(logOpts, zap.Error(err))...)
		return
	}

	// check if we want to log the slow query
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold && l.LogLevel >= gormlogger.Warn {
		logger.Warn("trace", logOpts...)
		return
	}

	// check if we want to log the query
	if l.LogLevel >= gormlogger.Info {
		logger.Debug("trace", logOpts...)
		return
	}
}

func (l CustomGormLogger) logger(ctx context.Context) *zap.Logger {
	logger := l.ZapLogger
	if l.Context != nil {
		fields := l.Context(ctx)
		logger = logger.With(fields...)
	}

	for i := 2; i < 15; i++ {
		_, file, _, ok := runtime.Caller(i)
		switch {
		case !ok:
		case strings.HasSuffix(file, "_test.go"):
		case strings.Contains(file, gormPackage):
		case strings.Contains(file, zapgormPackage):
		default:
			return logger.WithOptions(zap.AddCallerSkip(i))
		}
	}

	return logger
}
