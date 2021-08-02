package main

import (
	"wkgc/lib/core/config"
	"wkgc/lib/utils/dbt"
	"wkgc/lib/utils/iot"
)

func main() {
	// load config
	config.Load()
	// checkDataDir
	iotool := iot.NewIoTool()
	iotool.CheckAndCreateDir(config.Config.WorkDir)
	// Init LocalDatabase - Create Table
	dbt.InitLocalDatabase()
	// todo

}
