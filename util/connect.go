package util

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(dbName string, username string, password string) error {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 想要正确的处理 time.Time ，您需要带上 parseTime 参数
	dsn := "%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	dsn = fmt.Sprintf(dsn, username, password, dbName)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//GORM 使用 database/sql 维护连接池
	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	// 设置空闲连接的存活时间
	sqlDB.SetConnMaxIdleTime(time.Minute * 30)
	return err
}

func printLine() {
	fmt.Println("================")
}
