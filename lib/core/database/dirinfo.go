package database

import "gorm.io/gorm"

type DirInfo struct {
	gorm.Model
	Dirname  string
	Path     string
	Describe string
	Isgit    bool
	Lang     string
	Tags     []Tag `gorm:"many2many:dirinfo_tag"`
}

func (d *DirInfo) Init(dirname string, dirpath string, describe string, isgit bool, lang string) {
	d.Dirname = dirname
	d.Path = dirpath
	d.Describe = describe
	d.Isgit = isgit
	d.Lang = lang
}

func (d *DirInfo) AddTag(tag *Tag) {
	d.Tags = append(d.Tags, *tag)
}
func (d *DirInfo) AddTags(tag *[]Tag) {
	for i := 0; i < len(*tag); i++ {
		d.Tags = append(d.Tags, (*tag)[i])
	}
}
func (d *DirInfo) SetTags(tag *[]Tag) {
	d.Tags = *tag
}

/*
	Single insert
*/
func (d *DirInfo) Add() {
	DB.Create(d)
}

/*
	Single update
*/
func (d *DirInfo) Update() {
	DB.Save(d)
}

/*
	Single delete
*/
func (d *DirInfo) DeleteDirInfo() bool {
	if err := DB.Delete(d, d.ID); err != nil {
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
