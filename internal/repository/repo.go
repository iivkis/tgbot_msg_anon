package repository

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var (
	Users    = new(users)
	Messages = new(messages)
)

func init() {
	//init database
	{
		if database, err := gorm.Open(sqlite.Open("db.sqlite"), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
		}); err != nil {
			panic(err)
		} else {
			db = database
		}

		//create migrations
		if err := db.AutoMigrate(
			&UserModel{},
			&MessageModel{},
		); err != nil {
			panic(err)
		}
	}
}
