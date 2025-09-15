package recipes

type NewRecipeEvent struct {
	Recipe Recipe
}

func (e NewRecipeEvent) Name() string {
	return "NewRecipeEvent"
}
