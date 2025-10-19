package chores

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

	choreEventHandler := NewChoreEventHandler(repo)
	eventBus.Register("NewChoreEvent", choreEventHandler.handleChoreEvent)
	eventBus.Register("CompletedChoreEvent", choreEventHandler.handleChoreEvent)

	return &Module{
		Service: &service,
	}
}
