package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DirInfo struct {
	gorm.Model
	// gorm auto create id
	// Did      int    //`col:"did" json:"did"`
	Dirname  string //`col:"dirname" json:"dirname"`
	Path     string //
	Describe string //`col:"describe" json:"describe"`
	Isgit    bool   //`col:"isgit" json:"isgit"`
	Lang     string //`col:"lang" json:"lang"`
	Tag      string //`col:"tag" json:"tag"`
	// grom auto create about time
	// Created  string //`col:"created" json:"created"`
}

var db, _ = gorm.Open(sqlite.Open("data.db"), &gorm.Config{
	Logger: logger.Default.LogMode(logger.Silent),
})

func InitLocalDatabase() {
	// db := OpenLocalDatabases()
	if err := db.AutoMigrate(&DirInfo{}); err != nil {
		panic(err)
	}
}

func AddDirInfo(dirname string, dirpath string, describe string, isgit bool, lang string, tag string) {
	db.Create(&DirInfo{Dirname: dirname, Path: dirpath, Describe: describe, Isgit: isgit, Lang: lang, Tag: tag})
}

func (d *DirInfo) SelectDirinfoByDid(id int) bool {
	if err := db.First(d, id).Error; err != nil {
		return false
	}
	return true
}

func (d *DirInfo) SelectDirinfoByPath(path string) bool {
	if err := db.First(d, "path = ?", path).Error; err != nil {
		return false
	}
	return true
}

func (d *DirInfo) DeleteDirInfo() bool {
	if err := db.Delete(d, d.ID); err != nil {
		return false
	}
	return true
}
