package recipes

type NewRecipeDto struct {
	Title        string                `json:"title"`
	Username     string                `json:"author"`
	Description  string                `json:"description"`
	Ingredients  []RecipeIngredientDto `json:"ingredients"`
	Instructions []RecipeStepDto       `json:"instructions"`
}

type RecipeDto struct {
	Title        string                `json:"title"`
	Username     string                `json:"author"`
	Description  string                `json:"description"`
	Ingredients  []RecipeIngredientDto `json:"ingredients"`
	Instructions []RecipeStepDto       `json:"instructions"`
}

type RecipeIngredientDto struct {
	Ingredient string `json:"ingredient"`
	Amount     int    `json:"amount"`
	Unit       string `json:"unit"`
}

type RecipeStepDto struct {
	Text string `json:"text"`
	Time string `json:"time"`
}
