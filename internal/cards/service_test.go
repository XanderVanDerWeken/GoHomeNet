package cards

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
)

func TestGetAllCards(t *testing.T) {
	// Arrange
	mockRepo := newMockRepository()
	service := NewService(mockRepo)

	// Act
	cards, err := service.GetAllCards()

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if len(cards) != 2 {
		t.Errorf("expected 2 cards, got %d", len(cards))
	}
}

func TestCreateCard(t *testing.T) {
	// Arrange
	mockRepo := newMockRepository()
	service := NewService(mockRepo)

	createCardRequest := CreateCardRequest{
		Name:    "New Card",
		DueDate: time.Now(),
	}

	// Act
	card, err := service.CreateCard(createCardRequest)

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if card == nil {
		t.Error("expected a card, got nil")
	}
}

func TestDeleteCard(t *testing.T) {
	// Arrange
	mockRepo := newMockRepository()
	service := NewService(mockRepo)

	// Act
	err := service.DeleteCard(1)

	// Assert
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func newMockRepository() *mockRepository {
	return &mockRepository{}
}

type mockRepository struct {
	mock.Mock
}

func (m *mockRepository) GetAllCards() ([]Card, error) {
	return []Card{
		{Name: "Card 1", DueDate: time.Now()},
		{Name: "Card 2", DueDate: time.Now().Add(24 * time.Hour)},
	}, nil
}

func (m *mockRepository) CreateCard(card *Card) error {
	return nil
}

func (m *mockRepository) DeleteCard(id uint) error {
	return nil
}
