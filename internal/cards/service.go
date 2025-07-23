package cards

import (
	"errors"
	"log"

	errx "github.com/xandervanderweken/GoHomeNet/internal/errors"
	"gorm.io/gorm"
)

type Service interface {
	GetAllCards() ([]CardDto, error)
	CreateCard(request CreateCardRequest) (*CardDto, error)
	DeleteCard(id uint) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllCards() ([]CardDto, error) {
	log.Println("Fetching all cards from repository")

	cards, err := s.repo.GetAllCards()

	if err != nil {
		if errors.Is(err, gorm.ErrInvalidDB) {
			log.Println("Database connection error:", err)
			return nil, errx.ErrInternalServer
		}
	}

	cardDtos := make([]CardDto, 0, len(cards))
	for _, card := range cards {
		cardDtos = append(cardDtos, CardDto{
			ID:      card.ID,
			Name:    card.Name,
			DueDate: card.DueDate,
		})
	}
	return cardDtos, nil
}

func (s *service) CreateCard(request CreateCardRequest) (*CardDto, error) {
	log.Println("Adding a new card to the repository")

	card := &Card{
		Name:    request.Name,
		DueDate: request.DueDate,
	}

	if err := s.repo.CreateCard(card); err != nil {
		log.Println("Error creating card:", err)
		return nil, errx.ErrInternalServer
	}

	log.Println("Card created successfully:", card.ID)
	return &CardDto{
		ID:      card.ID,
		Name:    card.Name,
		DueDate: card.DueDate,
	}, nil
}

func (s *service) DeleteCard(id uint) error {
	log.Println("Deleting card with ID:", id)

	err := s.repo.DeleteCard(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		log.Println("Card not found for deletion:", id)
		return errx.ErrNotFound
	}

	return err
}
