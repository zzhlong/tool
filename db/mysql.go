package db

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

var Db *gorm.DB
var once sync.Once

type Option func(*sql.DB)

func MaxIdleConns(mc int) Option {
	return func(ci *sql.DB) {
		ci.SetMaxIdleConns(mc)
	}
}

func MaxOpenConns(mo int) Option {
	return func(ci *sql.DB) {
		ci.SetMaxOpenConns(mo)
	}
}

func ConnMaxLife(cml time.Duration) Option {
	return func(ci *sql.DB) {
		ci.SetConnMaxLifetime(cml)
	}
}

func NewMysqlDb(host string, port int, username, password, database string, options ...func(db *sql.DB)) {
	once.Do(func() {
		dsn := username + ":" + password + "@tcp(" + host + ":" + fmt.Sprintf("%d", port) + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // DSN data source name
			DefaultStringSize:         256,   // string 类型字段的默认长度
			DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
			SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
		}))
		if err != nil {
			panic(err)
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic(err)
		}
		for _, option := range options {
			option(sqlDB)
		}
		Db = db
	})
	return
}
