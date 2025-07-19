package chores

import "gorm.io/gorm"

type ChoreRepository interface {
	GetAllChores() ([]Chore, error)
	CreateChore(newChore *Chore) error
}

type choreRepository struct {
	db *gorm.DB
}

func NewChoreRepository(db *gorm.DB) ChoreRepository {
	return &choreRepository{db: db}
}

func (r *choreRepository) GetAllChores() ([]Chore, error) {
	var chores []Chore
	if err := r.db.Find(&chores).Error; err != nil {
		return nil, err
	}
	return chores, nil
}

func (r *choreRepository) CreateChore(newChore *Chore) error {
	return r.db.Create(newChore).Error
}
