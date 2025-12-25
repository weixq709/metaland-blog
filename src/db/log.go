package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wxq/metaland-blog/src/xzap"
	"go.uber.org/zap"
	glogger "gorm.io/gorm/logger"
)

const (
	infoStr      = "%s\n "
	warnStr      = "%s\n "
	traceStr     = "[%.3fms] [rows:%v] %s"
	traceWarnStr = "%s\n[%.3fms] [rows:%v] %s"
	traceErrStr  = "%s\n[%.3fms] [rows:%v] %s"
)

type Logger struct {
	logLevel      glogger.LogLevel
	SlowThreshold time.Duration
}

func (l *Logger) LogMode(level glogger.LogLevel) glogger.Interface {
	newLogger := *l
	newLogger.logLevel = level
	return &newLogger
}

func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= glogger.Info {
		xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().Infof(infoStr+msg, data...)
	}
}

func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logLevel >= glogger.Warn {
		xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().Warnf(warnStr+msg, data...)
	}
}

func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	// do nothing
}

func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)

	switch {
	case err != nil && l.logLevel >= glogger.Error && (!errors.Is(err, glogger.ErrRecordNotFound)):
		sql, rows := fc()
		if rows == -1 {
			xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().
				Errorf(traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().
				Errorf(traceErrStr, err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.logLevel >= glogger.Warn:
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		sql, rows := fc()
		if rows == -1 {
			xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().
				Warnf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().
				Warnf(traceWarnStr, slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case l.logLevel == glogger.Info:
		sql, rows := fc()
		if rows == -1 {
			xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().
				Infof(traceStr, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			xzap.GetLogger().WithOptions(zap.AddCallerSkip(2)).Sugar().
				Infof(traceStr, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
