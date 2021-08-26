package database

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Dirinfos []DirInfo `gorm:"many2many:dirinfo_tag"`
}

func (t *Tag) Create() {
	var tt Tag
	if tt.SelectTagByName(t.Name) {
		*t = tt
	} else {
		DB.Create(t)
	}
}

/*
	Single Save
*/
func (t *Tag) Save() {
	DB.Save(t)
}

/*
	Single delete
*/
func (t *Tag) DeleteTag() bool {
	if len(t.Dirinfos) > 0 {
		return false
	} else if err := DB.Delete(t, t.ID).Error; err != nil {
		return false
	} else {
		return true
	}
}

/*
	Batch add
*/
func AddTags(t *[]Tag) bool {
	var tt Tag
	for i := 0; i < len(*t); i++ {
		if tt.SelectTagByName((*t)[i].Name) {
			return false
		}
	}
	DB.CreateInBatches(t, len(*t))
	return true
}

/*
	Query by id
*/
func (t *Tag) SelectTagById(id int) bool {
	if err := DB.Preload("Dirinfos").First(t, id).Error; err != nil {
		return false
	}
	return true
}

/*
	Query by name
*/
func (t *Tag) SelectTagByName(name string) bool {
	if err := DB.Preload("Dirinfos").First(t, "name = ?", name).Error; err != nil {
		return false
	}
	return true
}
