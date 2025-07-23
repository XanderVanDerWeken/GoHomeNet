package users

import "gorm.io/gorm"

type Repository interface {
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (r *repo) FindByEmail(email string) (*User, error) {
	var u User
	err := r.db.Where("email = ?", email).First(&u).Error
	return &u, err
}

func (r *repo) FindByID(id uint) (*User, error) {
	var u User
	err := r.db.First(&u, id).Error
	return &u, err
}
