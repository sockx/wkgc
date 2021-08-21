package main

import (
	"strings"
	"wkgc/lib/core/config"
	"wkgc/lib/core/database"
	"wkgc/lib/utils/iot"
)

func main() {
	// load config
	config.Load()
	// checkDataDir
	iotool := iot.NewIoTool()
	iotool.CheckAndCreateDir(config.Config.WorkDir)
	// Init LocalDatabase - Create Table
	database.InitLocalDatabase()
	// scan local path
	dirs := iot.ScanDir(config.Config.WorkDir)

	var ds []database.DirInfo
	for i := 0; i < len(dirs); i++ {
		shortname := strings.ReplaceAll(dirs[i], config.Config.WorkDir, "")
		var d database.DirInfo
		if !d.SelectDirInfoByPath(shortname) {
			d.Init(shortname, shortname, shortname, iot.IsGit(dirs[i]), "")
			ds = append(ds, d)
		}
	}
	if len(ds) > 0 {
		database.DB.CreateInBatches(ds, len(ds))
	}

	println("Init Success!")
	// todo

}
