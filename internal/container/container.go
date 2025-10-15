package container

import (
	"github.com/xandervanderweken/GoHomeNet/internal/auth"
	"github.com/xandervanderweken/GoHomeNet/internal/cards"
	"github.com/xandervanderweken/GoHomeNet/internal/chores"
	"github.com/xandervanderweken/GoHomeNet/internal/config"
	"github.com/xandervanderweken/GoHomeNet/internal/database"
	"github.com/xandervanderweken/GoHomeNet/internal/events"
	"github.com/xandervanderweken/GoHomeNet/internal/finances"
	"github.com/xandervanderweken/GoHomeNet/internal/recipes"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
	"gorm.io/gorm"
)

type Container struct {
	DB *gorm.DB

	EventBus *events.EventBus

	UserSvc     users.Service
	AuthSvc     auth.Service
	CardsSvc    cards.Service
	ChoresSvc   chores.Service
	FinancesSvc finances.Service
	RecipeSvc   recipes.Service
}

func New() *Container {
	// Load Config
	config.Load("./config")

	// Connect to the database
	db := database.Connect()
	db.AutoMigrate(
		&users.User{},
		&cards.Card{},
		&chores.Chore{},
		&finances.Transaction{}, &finances.Category{},
		&recipes.Recipe{}, &recipes.RecipeIngredient{}, &recipes.RecipeStep{},
	)

	eventBus := events.NewEventBus()

	// Add Users Module
	userModule := users.RegisterModule(db)

	// Add Auth Module
	authModule := auth.RegisterModule(*userModule.Repo)

	// Add Cards Module
	cardModule := cards.RegisterModule(db, eventBus, *userModule.Repo)

	// Add Chores Module
	choreModule := chores.RegisterModule(db, eventBus, *userModule.Repo)

	// Add Finances Module
	financesModule := finances.RegisterModule(db, eventBus)

	// Add Recipe Module
	recipeModule := recipes.RegisterModule(db, eventBus, *userModule.Repo)

	return &Container{
		DB: db,

		EventBus: eventBus,

		UserSvc: *userModule.Service,
		AuthSvc: *authModule.Service,

		CardsSvc:    *cardModule.Service,
		ChoresSvc:   *choreModule.Service,
		FinancesSvc: *financesModule.Service,
		RecipeSvc:   *recipeModule.Service,
	}
}
