package cards

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	AddCard(userID uint, name string, expiresAt time.Time) error
	GetAllCards() []Card
	GetAllOwnCards(username string) ([]Card, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) AddCard(userID uint, name string, expiresAt time.Time) error {
	card := Card{
		UserID:    userID,
		Name:      name,
		ExpiresAt: expiresAt,
	}

	if err := r.db.Create(&card).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllCards() []Card {
	var cards []Card
	r.db.Find(&cards)

	return cards
}

func (r *repository) GetAllOwnCards(username string) ([]Card, error) {
	var cards []Card

	err := r.db.
		Joins("JOIN users ON users.id = cards.user_id").
		Where("users.username = ?", username).
		Find(&cards).Error

	return cards, err
}
