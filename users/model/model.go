package model

import (
	"go-micro-demo/users/common"
	 _ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"errors"
	"time"
	"github.com/micro/go-log"
)

var (
	Config *config
	l *dblog
)


type config struct {
	// users.mysqldsn root:root@/test?charset=utf8&parseTime=True&loc=Local
	mysqlDSN string
	db *gorm.DB
}

type dblog struct{}

func (d dblog)Print(v ...interface{})  {
	log.Log(v)
}

func Init() error {

	Config = new(config)

	l = new(dblog)

	kv := common.ConsulClient.KV()

	pair , _ , err := kv.Get("users.mysqldsn",nil)
	if err != nil {
		return err
	}

	Config.mysqlDSN = string(pair.Value)
	if Config.mysqlDSN == "" {
		return errors.New("mysql dns 配置为空")
	}

	// 链接mysql
	db, err := gorm.Open("mysql", Config.mysqlDSN)

	if err != nil {
		return err
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Hour)
	db.SetLogger(l)
	db.LogMode(true)
	db.Debug()
	db.DB().Ping()

	// 数据库迁移
	db.AutoMigrate(&User{})
	Config.db = db
	return nil
}

func NewDB() *gorm.DB {
	return Config.db.New()
}
