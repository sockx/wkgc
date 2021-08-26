package ctl

import (
	"wkgc/lib/core/database"
)

func GetAllDirifoList() []database.DirInfo {
	var dl []database.DirInfo
	database.DB.Preload("Tags").Find(&dl)
	return dl
}
