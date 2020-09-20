package store

import (
	"github.com/jinzhu/gorm"
	"joinc.co.kr/jwiki/model"
)

// WikiStore ...
type WikiStore struct {
	db *gorm.DB
}

// NewWikiStore ...
func NewWikiStore(db *gorm.DB) *WikiStore {
	return &WikiStore{db: db}
}

// GetByID ...
func (w *WikiStore) GetByID(id int) (*model.Wiki, error) {
	return &model.Wiki{Name: "yundream"}, nil
}

// SaveWikiPage ..
func (w *WikiStore) SaveWikiPage(wiki *model.Wiki) error {
	tx := w.db.Begin()
	if err := tx.Create(&wiki).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
