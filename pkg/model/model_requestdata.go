package model

import (
	"gorm.io/gorm"
)

type Requestdata struct {
	Model          `bson:"-"`
	MID_           string `bson:"_id" gorm:"-" json:"-" xml:"mid_"`
	Token          string `bson:"token" json:"token" xml:"token"`
	Operation      string `bson:"operation" json:"operation" xml:"operation"`
	Gitrepoowner   string `bson:"gitrepoowner" json:"gitrepoowner" xml:"gitrepoowner"`
	Repotodelete   string `bson:"repotodelete" json:"repotodelete" xml:"repotodelete"`
	Repotocommit   string `bson:"repotocommit" json:"repotocommit" xml:"repotocommit"`
	Branchtocommit string `bson:"branchtocommit" json:"branchtocommit" xml:"branchtocommit"`
}

func (Requestdata) TableName() string {
	return "requestdata"
}
func (m *Requestdata) PreloadRequestdata(db *gorm.DB) *gorm.DB {
	return db
}
