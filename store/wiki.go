package store

import (
	"joinc.co.kr/jwiki/model"
)

// WikiStore ...
type WikiStore struct {
	Name string
}

// NewWikiStore ...
func NewWikiStore(name string) *WikiStore {
	return &WikiStore{Name: name}
}

// GetByID ...
func (w *WikiStore) GetByID(id int) (*model.Wiki, error) {
	return &model.Wiki{Name: "yundream"}, nil
}
