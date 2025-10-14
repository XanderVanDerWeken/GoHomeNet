package finances

type CategoryRepository interface {
	AddCategory(newName string) error

	GetCategoryByName(name string) (*Category, error)
	GetAllCategories() ([]Category, error)
}
