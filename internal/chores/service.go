package chores

type ChoreService interface {
	GetAllChores() ([]Chore, error)
	CreateChore(request CreateChoreRequest) (*Chore, error)
}

type choreService struct {
	repo ChoreRepository
}

func NewChoreService(repo ChoreRepository) ChoreService {
	return &choreService{repo: repo}
}

func (s *choreService) GetAllChores() ([]Chore, error) {
	return s.repo.GetAllChores()
}

func (s *choreService) CreateChore(request CreateChoreRequest) (*Chore, error) {
	chore := &Chore{
		Name:        request.Name,
		Description: request.Description,
		DueDate:     request.DueDate,
		IsDone:      false,
	}

	if err := s.repo.CreateChore(chore); err != nil {
		return nil, err
	}

	return chore, nil
}
