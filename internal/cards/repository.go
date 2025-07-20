package cards

import "gorm.io/gorm"

type Repository interface {
	GetAllCards() ([]Card, error)
	CreateCard(card *Card) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllCards() ([]Card, error) {
	var cards []Card
	if err := r.db.Find(&cards).Error; err != nil {
		return nil, err
	}
	return cards, nil
}

func (r *repository) CreateCard(card *Card) error {
	return r.db.Create(card).Error
}
