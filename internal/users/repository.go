package users

import "gorm.io/gorm"

type Repository interface {
	SaveUser(newUser *User) error
	GetUserIdByUsername(username string) (uint, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByUserId(userId uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) SaveUser(newUser *User) error {
	if err := r.db.Create(newUser).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUserIdByUsername(username string) (uint, error) {
	user, err := r.GetUserByUsername(username)

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *repository) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return nil, ErrUserNotFound
	}

	return &user, nil
}

func (r *repository) GetUserByUserId(userId uint) (*User, error) {
	var user User
	if err := r.db.Find(&user, userId).Error; err != nil {
		return nil, ErrUserNotFound
	}

	return &user, nil
}
