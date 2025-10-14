package finances

type NewCategoryEvent struct {
	NewName string
}

func (e NewCategoryEvent) Name() string {
	return "NewCategoryEvent"
}
