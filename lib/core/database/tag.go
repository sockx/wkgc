package database

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name     string
	Dirinfos []DirInfo `gorm:"many2many:dirinfo_tag"`
}

func (t *Tag) Add() {
	DB.Create(t)
}

/*
	Single update
*/
func (t *Tag) Update() {
	DB.Save(t)
}

/*
	Single delete
*/
func (t *Tag) DeleteDirInfo() bool {
	if err := DB.Delete(t, t.ID); err != nil {
		return false
	}
	return true
}

/*
	Batch add
*/
func AddTag(t *[]Tag) {
	DB.CreateInBatches(t, len(*t))
}

/*
	Query by did
*/
func (t *Tag) SelectTagById(id int) bool {
	if err := DB.Preload("Dirinfos").First(t, id).Error; err != nil {
		return false
	}
	return true
}

/*
	Query by path
*/
func (t *Tag) SelectTagByPath(path string) bool {
	if err := DB.First(t, "path = ?", path).Error; err != nil {
		return false
	}
	return true
}
