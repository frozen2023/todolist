package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	USERNAME = "root"
	PASSWORD = "123456"
	URL      = "tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
)

var SqlSession *gorm.DB

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@%s", USERNAME, PASSWORD, URL)
	// 建立数据库连接
	var err error
	SqlSession, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败！")
	}
}
