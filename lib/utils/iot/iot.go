package iot

import (
	"io/ioutil"
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
	err := os.Mkdir(dirName, 0755)
	return err == nil
}

/*
Get the full path of all files in the current folder.
*/
func ScanDir(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("read [%s] error, err = %s\n", dir, err)
	}
	var fileList []string
	for _, file := range files {
		fileList = append(fileList, dir+file.Name())
	}
	return fileList
}

/**
Determine whether the current directory contains a .git folder to determine whether the current file is a git project.
*/
func IsGit(dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("read [%s] error, err = %s\n", dir, err)
	}
	for _, file := range files {
		if file.IsDir() && file.Name() == ".git" {
			return true
		}
	}
	return false
}
