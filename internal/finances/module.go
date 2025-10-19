package finances

import (
	"github.com/xandervanderweken/GoHomeNet/internal/events"
	"gorm.io/gorm"
)

type Module struct {
	Service *Service
}

func RegisterModule(db *gorm.DB, eventBus *events.EventBus) *Module {
	categoryRepo := NewCategoryRepository(db)
	transactionRepo := NewTransactionRepository(db)

	service := NewService(transactionRepo, categoryRepo, eventBus)

	return &Module{
		Service: &service,
	}
}
