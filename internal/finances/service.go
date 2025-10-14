package finances

import "github.com/xandervanderweken/GoHomeNet/internal/events"

type CategoryService interface {
	AddCategory(newName string) error
	GetAllCategories() ([]Category, error)
}

type categoryService struct {
	repo     CategoryRepository
	eventBus *events.EventBus
}

func NewCategoryService(repo CategoryRepository, eventBus *events.EventBus) CategoryService {
	return &categoryService{repo: repo, eventBus: eventBus}
}

func (s *categoryService) AddCategory(newName string) error {
	_, err := s.repo.GetCategoryByName(newName)
	if err != nil {
		return err
	}

	s.eventBus.Publish(NewCategoryEvent{
		newName,
	})

	return nil
}

func (s *categoryService) GetAllCategories() ([]Category, error) {
	return s.repo.GetAllCategories()
}
