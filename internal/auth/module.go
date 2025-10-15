package auth

import (
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type Module struct {
	Service *Service
}

func RegisterModule(userRepo users.Repository) *Module {
	service := NewService(userRepo)

	return &Module{
		Service: &service,
	}
}
