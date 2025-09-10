package users

import "gorm.io/gorm"

type Repository interface {
	SaveUser(username, password, firstName, lastName string) error
	GetUserIdByUsername(username string) (uint, error)
	GetUserByUserId(userId uint) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) SaveUser(username, password, firstName, lastName string) error {
	user := User{
		Username:  username,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *repository) GetUserIdByUsername(username string) (uint, error) {
	var user User
	err := r.db.Where("username = ?", username).First(&user).Error

	if err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (r *repository) GetUserByUserId(userId uint) (*User, error) {
	var user User
	if err := r.db.Find(&user, userId).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
