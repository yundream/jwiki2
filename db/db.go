package db

import (
	"github.com/jinzhu/gorm"

	// mysql
	_ "github.com/go-sql-driver/mysql"

	"joinc.co.kr/jwiki/model"
)

// New ...
func New() (*gorm.DB, error) {
	db, err := gorm.Open("mysql", "root:1234@tcp(172.17.0.2:3306)/wiki")
	if err != nil {
		return nil, err
	}

	db.DB().SetMaxIdleConns(5)
	db.LogMode(true)
	return db, nil
}

// AutoMigration ...
func AutoMigration(db *gorm.DB) {
	db.AutoMigrate(
		&model.Wiki{},
	)
}
