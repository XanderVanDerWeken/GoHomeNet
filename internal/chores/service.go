package chores

import (
	"errors"
	"log"

	errx "github.com/xandervanderweken/GoHomeNet/internal/errors"
	"gorm.io/gorm"
)

type Service interface {
	GetAllChores() ([]ChoreDto, error)
	CreateChore(request CreateChoreRequest) (*ChoreDto, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllChores() ([]ChoreDto, error) {
	log.Println("Fetching all chores from repository")

	chores, err := s.repo.GetAllChores()

	if err != nil {
		if errors.Is(err, gorm.ErrInvalidDB) {
			log.Println("Database connection error:", err)
			return nil, errx.ErrInternalServer
		}
	}

	choreDtos := make([]ChoreDto, 0, len(chores))
	for _, chore := range chores {
		choreDtos = append(choreDtos, ChoreDto{
			ID:          chore.ID,
			Name:        chore.Name,
			Description: chore.Description,
			DueDate:     chore.DueDate,
			IsDone:      chore.IsDone,
		})
	}
	return choreDtos, nil
}

func (s *service) CreateChore(request CreateChoreRequest) (*ChoreDto, error) {
	log.Println("Adding a new chore to the repository")

	chore := &Chore{
		Name:        request.Name,
		Description: request.Description,
		DueDate:     request.DueDate,
		IsDone:      false,
	}

	if err := s.repo.CreateChore(chore); err != nil {
		log.Println("Error creating chore:", err)
		return nil, errx.ErrInternalServer
	}

	log.Println("Chore created successfully:", chore.ID)
	return &ChoreDto{
		ID:          chore.ID,
		Name:        chore.Name,
		Description: chore.Description,
		DueDate:     chore.DueDate,
		IsDone:      chore.IsDone,
	}, nil
}
