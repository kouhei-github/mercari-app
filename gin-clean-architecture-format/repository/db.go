package repository

import (
	"kouhei-github/sample-gin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func init() {
	// connect DB
	conf := config.Config
	dsn := conf.DbUser + ":" + conf.DbPassword + "@tcp(" + conf.DbHost + ":" + conf.DbPort + ")/" + conf.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:root@tcp(127.0.0.1:8889)/app2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&BlogEntity{})
}
