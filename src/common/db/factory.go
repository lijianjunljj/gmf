package db

import (
	"fmt"
	"gmf/src/common/config"
	"gorm.io/gorm"
)

var DB AbstractDatabase

func Init(dbType string, configFunc func() interface{}) {
	fmt.Println("dbType", dbType)
	if dbType == "mysql" {
		conf := configFunc()
		fmt.Println("conf", conf)
		DB = NewMysql(false, conf.(*config.MysqlOptions))
	}
}

func GetDB() *gorm.DB {
	return DB.DB()
}
