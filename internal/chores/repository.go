package chores

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateChore(newChore *Chore) error
	GetAllChores() []Chore
	GetChoresByUsername(username string) ([]Chore, error)
	GetChoreById(choreID uint) (*Chore, error)
	CompleteChore(choreID uint) error
	DeleteChore(choreID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateChore(newChore *Chore) error {
	if err := r.db.Create(newChore).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetAllChores() []Chore {
	var chores []Chore

	r.db.Find(&chores)

	return chores
}

func (r *repository) GetChoresByUsername(username string) ([]Chore, error) {
	var chores []Chore

	err := r.db.
		Joins("JOIN users ON users.id = chores.user_id").
		Where("users.username = ?", username).
		Find(&chores).Error

	return chores, err
}

func (r *repository) GetChoreById(choreID uint) (*Chore, error) {
	var chore Chore
	if err := r.db.First(&chore, choreID).Error; err != nil {
		return nil, err
	}

	return &chore, nil
}

func (r *repository) CompleteChore(choreID uint) error {
	return r.db.Model(&Chore{}).
		Where("id = ?", choreID).
		Update("completed", true).Error
}

func (r *repository) DeleteChore(choreID uint) error {
	return r.db.Delete(&Chore{}, choreID).Error
}
