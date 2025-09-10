package chores

import (
	"time"

	"gorm.io/gorm"
)

type Repsitory interface {
	CreateChore(userId uint, title, notes string, dueDate *time.Time) error
	GetAllChores() []Chore
	GetChoresByUsername(username string) ([]Chore, error)
	CompleteChore(choreID uint) error
	DeleteChore(choreID uint) error
}

type repsitory struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repsitory {
	return &repsitory{db: db}
}

func (r *repsitory) CreateChore(userId uint, title, notes string, dueDate *time.Time) error {
	chore := Chore{
		UserID:  userId,
		Title:   title,
		Notes:   notes,
		DueDate: dueDate,
	}

	if err := r.db.Create(&chore).Error; err != nil {
		return err
	}

	return nil
}

func (r *repsitory) GetAllChores() []Chore {
	var chores []Chore

	r.db.Find(&chores)

	return chores
}

func (r *repsitory) GetChoresByUsername(username string) ([]Chore, error) {
	var chores []Chore

	err := r.db.
		Joins("JOIN users ON users.id = chores.user_id").
		Where("users.username = ?", username).
		Find(&chores).Error

	return chores, err
}

func (r *repsitory) CompleteChore(choreID uint) error {
	return r.db.Model(&Chore{}).
		Where("id = ?", choreID).
		Update("completed", true).Error
}

func (r *repsitory) DeleteChore(choreID uint) error {
	return r.db.Delete(&Chore{}, choreID).Error
}
