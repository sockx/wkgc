package main

import (
	"strings"
	"testing"
	"wkgc/lib/core/config"
	"wkgc/lib/core/database"
	"wkgc/lib/utils/iot"
)

func Test_main(t *testing.T) {
	// load config
	config.Load()

	println(config.Config.WorkDir)

	// checkDataDir
	iotool := iot.NewIoTool()
	iotool.CheckAndCreateDir(config.Config.WorkDir)
	// Init LocalDatabase - Create Table
	database.InitLocalDatabase()
	// scan local path
	dirs := iot.ScanDir(config.Config.WorkDir)

	var ds []database.DirInfo
	var tag database.Tag
	tag.Name = "Test Init Tag"
	tag.Create()
	for i := 0; i < len(dirs); i++ {
		shortname := strings.ReplaceAll(dirs[i], config.Config.WorkDir, "")
		var d database.DirInfo
		if !d.SelectDirInfoByPath(shortname) {
			d.Init(shortname, shortname, shortname, iot.IsGit(dirs[i]), "")
			d.AddTag(&tag)
			ds = append(ds, d)
		}
	}
	if len(ds) > 0 {
		database.DB.CreateInBatches(ds, len(ds))
	}

	println("Init Success!")
	// todo

}

func Test_localDB(t *testing.T) {
	config.Load()
	println(config.Config.Database)
}
