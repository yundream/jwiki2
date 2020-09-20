package model

import (
	"github.com/jinzhu/gorm"
)

// Wiki ...
type Wiki struct {
	gorm.Model
	Name     string `gorm:"column:name;size:160"`
	Title    string `gorm:"column:title;size:160"`
	Author   string `gorom:"size:80"`
	Contents string `gorm:"column:contents"`
}
