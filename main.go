package main

import (
	"wkgc/lib/core/config"
	"wkgc/lib/utils/iot"
)

func main() {
	// load config
	config.Load()
	// checkDataDir
	iotool := iot.NewIoTool()
	iotool.CheckAndCreateDir(config.Config.WorkDir)
	// todo

}
