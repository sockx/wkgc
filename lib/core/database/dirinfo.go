package database

import (
	"gorm.io/gorm"
)

type DirInfo struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey" json:"id"`
	Dirname  string `json:"dirname"`
	Path     string `json:"path"`
	Describe string `json:"describe"`
	Isgit    bool   `json:"isgit"`
	Lang     string `json:"lang"`
	Tags     []Tag  `gorm:"many2many:dirinfo_tag" json:"tags"`
}

func (d *DirInfo) Init(dirname string, dirpath string, describe string, isgit bool, lang string) {
	d.Dirname = dirname
	d.Path = dirpath
	d.Describe = describe
	d.Isgit = isgit
	d.Lang = lang
}

func (d *DirInfo) AddTag(tag *Tag) {
	var flag = false
	for i := 0; i < len(d.Tags); i++ {
		if d.Tags[i].Name == (*tag).Name {
			flag = true
		}
	}
	if !flag {
		d.Tags = append(d.Tags, *tag)
	}
}
func (d *DirInfo) AddTags(tag *[]Tag) {
	flag := false
	for i := 0; i < len(*tag); i++ {
		flag = false
		for j := 0; j < len(d.Tags); j++ {
			if d.Tags[j].Name == (*tag)[i].Name {
				flag = true
			}
		}
		if !flag {
			d.Tags = append(d.Tags, (*tag)[i])
		}
	}
}
func (d *DirInfo) SetTags(tag *[]Tag) {
	d.Tags = *tag
}

/*
	Single insert
*/
func (d *DirInfo) Create() {
	var td DirInfo
	if td.SelectDirInfoByPath(d.Path) {
		td.AddTags(&d.Tags)
		*d = td
		d.Save()
	} else {
		DB.Create(d)
	}
}

/*
	Single update
*/
func (d *DirInfo) Save() {
	DB.Save(d)
}

/*
	Single delete
*/
func (d *DirInfo) DeleteDirInfo() bool {
	if err := DB.Delete(d, d.ID).Error; err != nil {
		return false
	}
	return true
}

/*
	Batch add
*/
func AddDirInfo(d *[]DirInfo) {
	DB.CreateInBatches(d, len(*d))
}

/*
	Query by did
*/
func (d *DirInfo) SelectDirInfoByDid(id int) bool {
	if err := DB.Preload("Tags").First(d, id).Error; err != nil {
		return false
	}
	return true
}

/*
	Query by path
*/
func (d *DirInfo) SelectDirInfoByPath(path string) bool {
	if err := DB.First(d, "path = ?", path).Error; err != nil {
		return false
	}
	return true
}
