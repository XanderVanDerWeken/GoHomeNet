package finances

type NewCategoryEvent struct {
	NewCategory Category
}

func (e NewCategoryEvent) Name() string {
	return "NewCategoryEvent"
}

type NewTransactionEvent struct {
	NewTransaction Transaction
}

func (e NewTransactionEvent) Name() string {
	return "NewTransactionEvent"
}
