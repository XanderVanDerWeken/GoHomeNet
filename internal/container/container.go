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

	UserRepo users.Repository
	UserSvc  users.Service

	AuthSvc auth.Service

	CardsRepo cards.Repository
	CardsSvc  cards.Service

	ChoresRepo chores.Repository
	ChoresSvc  chores.Service

	TransactionRepo finances.TransactionRepository
	CategoryRepo    finances.CategoryRepository
	FinancesSvc     finances.Service

	RecipeRepo recipes.Repository
	RecipeSvc  recipes.Service
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
	userRepo := users.NewRepository(db)
	userSvc := users.NewService(userRepo)

	// Add Auth Module
	authSvc := auth.NewService(userRepo)

	// Add Cards Module
	cardRepo := cards.NewRepository(db)
	cardSvc := cards.NewService(cardRepo, userRepo, eventBus)

	// Add Chores Module
	choreRepo := chores.NewRepository(db)
	choreSvc := chores.NewService(choreRepo, userRepo, eventBus)

	// Add Finances Module
	transactionRepo := finances.NewTransactionRepository(db)
	categoryRepo := finances.NewCategoryRepository(db)
	financesSvc := finances.NewService(transactionRepo, categoryRepo, eventBus)

	// Add Recipe Module
	recipeRepo := recipes.NewRepository(db)
	recipeSvc := recipes.NewService(recipeRepo, userRepo, eventBus)

	return &Container{
		DB: db,

		EventBus: eventBus,

		UserRepo: userRepo,
		UserSvc:  userSvc,

		AuthSvc: authSvc,

		CardsRepo: cardRepo,
		CardsSvc:  cardSvc,

		ChoresRepo: choreRepo,
		ChoresSvc:  choreSvc,

		TransactionRepo: transactionRepo,
		CategoryRepo:    categoryRepo,
		FinancesSvc:     financesSvc,

		RecipeRepo: recipeRepo,
		RecipeSvc:  recipeSvc,
	}
}
