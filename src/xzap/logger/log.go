package logger

import (
	"github.com/wxq/metaland-blog/src/xzap"
	"go.uber.org/zap"
)

func Debug(msg string, fields ...zap.Field) {
	xzap.GetLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	xzap.GetLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	xzap.GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	xzap.GetLogger().Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	xzap.GetLogger().Panic(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	xzap.GetLogger().Sugar().Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	xzap.GetLogger().Sugar().Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	xzap.GetLogger().Sugar().Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	xzap.GetLogger().Sugar().Errorf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	xzap.GetLogger().Sugar().Panicf(template, args...)
}
