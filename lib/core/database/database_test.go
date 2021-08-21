package database

import (
	"testing"
)

func TestInit(t *testing.T) {
	InitLocalDatabase()
}

func TestAdd(t *testing.T) {
	var d DirInfo
	var tag Tag
	tag.Name = "test"
	d.AddTag(&tag)
	d.Init("admin", "test/admin", "this is a test", false, "golang")
	d.Add()
	t.Log("add test")
}

func TestSelect(t *testing.T) {
	var d1 DirInfo
	if d1.SelectDirInfoByDid(1) {
		println(d1.ID)
		if d1.Tags != nil {
			for i := 0; i < len(d1.Tags); i++ {
				println(d1.Tags[i].Name)
			}
		}

		var t1 Tag
		if t1.SelectTagById(1) {
			if t1.Dirinfos != nil {
				for i := 0; i < len(t1.Dirinfos); i++ {
					println(t1.Dirinfos[i].Path)
				}
			}
		}

	} else {
		println("no id 1")
	}
	if d1.SelectDirInfoByDid(2) {
		println(d1.ID)
	} else {
		println("no id 2")
	}
}

func TestDelete(t *testing.T) {
	var d1 DirInfo
	d1.SelectDirInfoByDid(1)
	d1.DeleteDirInfo()
}
