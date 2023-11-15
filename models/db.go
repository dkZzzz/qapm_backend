package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:20021210@(127.0.0.1:13306)/qapm?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	DB = db
	db.AutoMigrate(&Message{}, &Compare{}, &Pa{}, &Optimize{})
	return db, err
}
