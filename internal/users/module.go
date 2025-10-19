package users

import (
	"gorm.io/gorm"
)

type Module struct {
	Service *Service
	Repo    *Repository // TODO: Check if makes sense to expose repo
}

func RegisterModule(db *gorm.DB) *Module {
	repo := NewRepository(db)
	service := NewService(repo)

	return &Module{
		Service: &service,
		Repo:    &repo,
	}
}
