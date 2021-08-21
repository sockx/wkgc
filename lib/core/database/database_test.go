package database

import (
	"fmt"
	"testing"
)

func TestInit(t *testing.T) {
	InitLocalDatabase()
}

func TestAdd(t *testing.T) {
	AddDirInfo("admin", "test/admin", "this is a test", false, "golang", "gotag")
	t.Log("add test")
}

func TestSelect(t *testing.T) {
	var d1 DirInfo
	if d1.SelectDirInfoByDid(1) {
		fmt.Print(d1.ID)
	} else {
		println("no")
	}
	if d1.SelectDirInfoByDid(2) {
		fmt.Print(d1.ID)
	} else {
		println("no")
	}
}

func TestDelete(t *testing.T) {
	var d1 DirInfo
	d1.SelectDirInfoByDid(2)
	d1.DeleteDirInfo()
}
