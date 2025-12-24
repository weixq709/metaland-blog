package xzap

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	logFileExtension = ".log"
	// 单个日志最大字节数（兆字节）
	maxSize    = 10
	maxBackups = 3
)

var logger *zap.Logger

type LogConfig struct {
	Level    string
	Path     string
	FileName string
	KeepDays int
}

func Initialize(cfg *viper.Viper) {
	config := defaultConfig()
	if err := cfg.Unmarshal(config); err != nil {
		panic(err)
	}

	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		// TODO 打印日志级别配置错误提示
		// 使用默认日志级别
		level = zapcore.InfoLevel
	}

	cores := make([]zapcore.Core, 0, 3)

	defaultFile := filepath.Join(config.Path, generateFileName(config.FileName, ""))
	errFile := filepath.Join(config.Path, generateFileName(config.FileName, "error"))

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		if level <= zapcore.WarnLevel {
			return level <= lvl && lvl <= zapcore.WarnLevel
		}
		return lvl <= zapcore.WarnLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(getEncoderConfig()), getLogWriter(defaultFile, config.KeepDays), highPriority))
	cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(getEncoderConfig()), getLogWriter(errFile, config.KeepDays), lowPriority))
	cores = append(cores, getConsoleCore(level))
	logger = zap.New(zapcore.NewTee(cores...), zap.AddCaller(), zap.AddCallerSkip(1))
	logger.Sugar().Infof("init logger success")
}

func GetLogger() *zap.Logger {
	return logger
}

func getConsoleCore(level zapcore.Level) zapcore.Core {
	levelEnabler := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})
	encoderConfig := getEncoderConfig()
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.Lock(os.Stdout), levelEnabler)
}

func getLogWriter(fileName string, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoderConfig() zapcore.EncoderConfig {
	config := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "name",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalLevelEncoder,
		EncodeTime: func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
			encoder.AppendString(time.Format("2006-01-02 15:04:05"))
		},
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
	return config
}

func defaultConfig() *LogConfig {
	return &LogConfig{
		Level:    "info",
		Path:     "./logs",
		FileName: "app.log",
		KeepDays: 7,
	}
}

// generateFileName 生成指定后缀日志文件名称
//
// eg:
//
//	generateFileName("app.log") => app.log
//	generateFileName("app.log") => app-error.log
func generateFileName(fileName string, suffix string) string {
	var extension string
	var simpleFileName string
	res := strings.Split(fileName, ".")
	if len(res) == 0 {
		return ""
	}
	simpleFileName = res[0]

	if len(suffix) > 0 {
		simpleFileName = fmt.Sprintf("%s%s", simpleFileName, suffix)
	}

	if len(res) > 1 {
		extension = res[1]
	} else {
		extension = logFileExtension
	}
	return fmt.Sprintf("%s.%s", simpleFileName, extension)
}
