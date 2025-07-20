package cards

import "log"

type Service interface {
	CreateCard(request CreateCardRequest) (*CardDto, error)
	GetAllCards() ([]CardDto, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateCard(request CreateCardRequest) (*CardDto, error) {
	card := &Card{
		Name:    request.Name,
		DueDate: request.DueDate,
	}

	if err := s.repo.CreateCard(card); err != nil {
		return nil, err
	}

	return &CardDto{
		ID:      card.ID,
		Name:    card.Name,
		DueDate: card.DueDate,
	}, nil
}

func (s *service) GetAllCards() ([]CardDto, error) {
	log.Println("Fetching all cards from repository")

	cards, err := s.repo.GetAllCards()

	if err != nil {
		return nil, err
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
