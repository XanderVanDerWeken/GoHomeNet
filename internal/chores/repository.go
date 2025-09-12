package chores

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateChore(userId uint, title, notes string, dueDate *time.Time) error
	GetAllChores() []Chore
	GetChoresByUsername(username string) ([]Chore, error)
	CompleteChore(choreID uint) error
	DeleteChore(choreID uint) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateChore(userId uint, title, notes string, dueDate *time.Time) error {
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

func (r *repository) CompleteChore(choreID uint) error {
	return r.db.Model(&Chore{}).
		Where("id = ?", choreID).
		Update("completed", true).Error
}

func (r *repository) DeleteChore(choreID uint) error {
	return r.db.Delete(&Chore{}, choreID).Error
}
