package users

import "gorm.io/gorm"

type Repository interface {
	SaveUser(username, password, firstName, lastName string) error
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
