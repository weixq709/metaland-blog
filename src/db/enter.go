package db

import (
	"errors"
	"fmt"
	"time"

	"github.com/spf13/viper"
	xzap "github.com/wxq/metaland-blog/src/xzap/logger"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var GlobalDB *gorm.DB

type DbConfig struct {
	Username      string      `mapstructure:"username"`
	Password      string      `mapstructure:"password"`
	Host          string      `mapstructure:"host"`
	Port          int         `mapstructure:"port"`
	Schema        string      `mapstructure:"schema"`
	LogLevel      string      `mapstructure:"log_level"`
	SlowThreshold int         `mapstructure:"slow_threshold"`
	Pool          *PoolConfig `mapstructure:"pool"`
}

type PoolConfig struct {
	MaxOpenConnection int `mapstructure:"max_open_connection"`
	MaxIdleConnection int `mapstructure:"max_idle_connection"`
	ConnMaxLifetime   int `mapstructure:"conn_max_lifetime"`
}

func Initialize(cfg *viper.Viper) error {
	config := defaultConfig()
	if err := cfg.Sub("db").Unmarshal(config); err != nil {
		return err
	}
	if err := config.Check(); err != nil {
		return err
	}
	if config.Pool == nil {
		config.Pool = defaultPoolConfig()
	}

	dsn := getDsn(config)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: &Logger{logLevel: glogger.Info, SlowThreshold: time.Duration(config.SlowThreshold) * time.Second},
	})
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to connect database, error: %v", err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to initialize the connection pool, error: %v", err))
	}

	sqlDB.SetMaxOpenConns(config.Pool.MaxOpenConnection)
	sqlDB.SetMaxIdleConns(config.Pool.MaxIdleConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Pool.ConnMaxLifetime) * time.Second)

	if sqlDB.Ping() != nil {
		// 连接超时
		return errors.New(fmt.Sprintf("Failed to ping database, error: %v", err))
	}

	GlobalDB = db
	xzap.Infof("Initialize datasource successfully")
	return nil
}

func (cfg *DbConfig) Check() error {
	if cfg.Username == "" {
		return errors.New("username is required")
	}
	if cfg.Password == "" {
		return errors.New("password is required")
	}
	if cfg.Schema == "" {
		return errors.New("schema is required")
	}
	return nil
}

func defaultConfig() *DbConfig {
	return &DbConfig{
		Host:          "localhost",
		Port:          3306,
		LogLevel:      "info",
		SlowThreshold: 5,
	}
}

func defaultPoolConfig() *PoolConfig {
	return &PoolConfig{
		MaxOpenConnection: 10,
		MaxIdleConnection: 10,
		ConnMaxLifetime:   60,
	}
}

func getDsn(config *DbConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Schema)
}
