package cards

import (
	"github.com/xandervanderweken/GoHomeNet/internal/events"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
	"gorm.io/gorm"
)

type Module struct {
	Service *Service
}

func RegisterModule(db *gorm.DB, eventBus *events.EventBus, userRepo users.Repository) *Module {
	repo := NewRepository(db)
	service := NewService(repo, userRepo, eventBus)

	return &Module{
		Service: &service,
	}
}
