package db

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var GlobalDB *gorm.DB

func Initialize() {
	dsn := "root:123456@tcp(192.168.3.124:3306)/metaland-blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,
		DisableDatetimePrecision: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to connect database, error: %v", err)))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(errors.New(fmt.Sprintf("Failed to initialize the connection pool, error: %v", err)))
	}

	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	GlobalDB = db
}
