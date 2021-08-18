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
	for i := 0; i < len(dirs); i++ {
		// println(dirs[i], iot.IsGit(dirs[i]))
		shortname := strings.ReplaceAll(dirs[i], config.Config.WorkDir, "")
		var d database.DirInfo
		if !d.SelectDirinfoByPath(shortname) {
			database.AddDirInfo(shortname, shortname, shortname, iot.IsGit(dirs[i]), "", "")
		}
	}

	println("Init Success!")
	// todo

}
