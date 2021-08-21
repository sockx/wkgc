package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB, _ = gorm.Open(sqlite.Open("data.db"), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Silent),
})

func InitLocalDatabase() {
	// db := OpenLocalDatabases()
	if err := DB.AutoMigrate(&DirInfo{}, &Tag{}); err != nil {
		panic(err)
	}
}
