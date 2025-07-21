package chores

import "gorm.io/gorm"

type Repository interface {
	GetAllChores() ([]Chore, error)
	CreateChore(newChore *Chore) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAllChores() ([]Chore, error) {
	var chores []Chore
	if err := r.db.Find(&chores).Error; err != nil {
		return nil, err
	}
	return chores, nil
}

func (r *repository) CreateChore(newChore *Chore) error {
	return r.db.Create(newChore).Error
}
