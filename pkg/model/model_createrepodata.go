package model

import (
	"gorm.io/gorm"
)

type Createrepodata struct {
	Model             `bson:"-"`
	MID_              string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Name              string `bson:"name" json:"name" xml:"name"`
	Autoinit          bool   `bson:"auto_init" json:"auto_init" xml:"auto_init"`
	Private           bool   `bson:"private" json:"private" xml:"private"`
	Gitignoretemplate string `bson:"gitignore_template" json:"gitignore_template" xml:"gitignore_template"`
}

func (Createrepodata) TableName() string {
	return "createrepodata"
}
func (m *Createrepodata) PreloadCreaterepodata(db *gorm.DB) *gorm.DB {
	return db
}
