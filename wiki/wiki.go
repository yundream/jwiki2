package wiki

import (
	"joinc.co.kr/jwiki/model"
)

// Store ...
type Store interface {
	GetByID(int) (*model.Wiki, error)
}
