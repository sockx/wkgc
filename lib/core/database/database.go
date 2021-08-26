package database

import (
	"wkgc/lib/core/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitLocalDatabase() {
	DB, _ = gorm.Open(sqlite.Open(config.Config.Database), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	// db := OpenLocalDatabases()
	if err := DB.AutoMigrate(&DirInfo{}, &Tag{}); err != nil {
		panic(err)
	}
}
