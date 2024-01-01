package models

import (
	"github.com/dkZzzz/qapm_backend/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitDB() (*gorm.DB, error) {
	user := config.DbUser
	password := config.DbPassword
	str := user + ":" + password + "@(" + config.DbHost + ":" + config.DbPort + ")/" + config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", str)
	if err != nil {
		return nil, err
	}
	DB = db
	db.AutoMigrate(&Message{}, &Compare{}, &Pa{}, &Optimize{})
	return db, err
}
