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
	tag.Create()

	println("tagid -> ", tag.ID)
	println("tagname -> ", tag.Name)

	d.AddTag(&tag)
	d.Init("admin", "test/admin", "this is a test", false, "golang")
	d.Create()
	println("d.id -> ", d.ID)
	println("d.path -> ", d.Path)
}

func TestSelect(t *testing.T) {
	var d1 DirInfo
	if d1.SelectDirInfoByDid(1) {
		println(d1.ID)
		if d1.Tags != nil {
			for i := 0; i < len(d1.Tags); i++ {
				println("d1 -> Tag Name", d1.Tags[i].Name)
			}
		}

		var t1 Tag
		if t1.SelectTagById(1) {
			if t1.Dirinfos != nil {
				for i := 0; i < len(t1.Dirinfos); i++ {
					println("t1 -> Dirinfos Path", t1.Dirinfos[i].Path)
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

func TestDeleteTag(t *testing.T) {
	var t1 Tag
	if t1.SelectTagByName("test") {
		res := t1.DeleteTag()
		println(res)
	}
}

func TestDelete(t *testing.T) {
	var d1 DirInfo
	if d1.SelectDirInfoByPath("test/admin") {
		if d1.DeleteDirInfo() {
			println("delete d1")
		}
	}

	var t1 Tag
	if t1.SelectTagByName("test") {
		println(t1.Name)
		println(len(t1.Dirinfos))
		println("tag -> Dirinfos", t1.Dirinfos)
	} else {
		println("no t1")
	}
}
