package db

import "gorm.io/gorm"

type AbstractDatabase interface {
	DB() *gorm.DB
	Connect() *gorm.DB
	AutoMigrate(db *gorm.DB, dst ...interface{})
}
