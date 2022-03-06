package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

var (
	TagNormalState = 1
	TagDelState    = 0
)

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
func TagIsExist(name string) bool {
	var tag Tag
	db.Where("name = ?", name).Select("id").First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
func AddTag(name string, createdBy string, state int) bool {
	maps := map[string]interface{}{
		"name":      name,
		"createdBy": createdBy,
		"state":     state,
	}
	a := maps["name"].(int)
	fmt.Println(a)
	return true
	//db.Create(&Tag{
	//	Name:      maps["name"].(string),
	//	State:     state,
	//	CreatedBy: createdBy,
	//})
	//return true
}
func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	err := scope.SetColumn("ModifiedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}
