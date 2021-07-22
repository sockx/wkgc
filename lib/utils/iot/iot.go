package iot

import (
	"log"
	"os"
)

type IoToolModel struct {
	NoWorkDir string
}

func NewIoTool() *IoToolModel {
	fpath, err := os.Getwd()
	if err != nil {
		log.Fatalln("Get Pwd err:", err)
	}
	return &IoToolModel{
		NoWorkDir: fpath,
	}
}

func (i *IoToolModel) CheckAndCreateDir(fpath string) {
	if !i.IsDir(fpath) {
		if !i.createDir(fpath) {
			log.Fatalln("目录创建失败！")
		}
	}
}

/*
判断文件是否存在
*/
func (i *IoToolModel) IsDir(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		// log.Println(err)
		return false
	}
	return s.IsDir()
}

/*
创建文件夹
*/
func (i *IoToolModel) createDir(dirName string) bool {
	err := os.Mkdir(dirName, 755)
	if err != nil {
		// log.Println(err)
		return false
	}
	return true
}
