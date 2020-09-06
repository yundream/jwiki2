package model

import (
	"github.com/jinzhu/gorm"
)

// Wiki ...
type Wiki struct {
	gorm.Model
	Name     string `gorm:"column:name;size:160" json:"name"`
	Title    string `gorm:"column:title;size:160" json:"title"`
	Author   string `gorom:"size:80" json:"author"`
	Contents string `gorm:"column:contents" json:"contents"`
}
