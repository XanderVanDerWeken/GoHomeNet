package recipes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/recipes"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type RecipeHandler struct {
	service     recipes.Service
	userService users.Service
}

func NewRecipeHandler(service recipes.Service, userSerivce users.Service) *RecipeHandler {
	return &RecipeHandler{
		service:     service,
		userService: userSerivce,
	}
}

func (h *RecipeHandler) PostNewRecipe(w http.ResponseWriter, r *http.Request) {
	var dto NewRecipeDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteError(w, err)
		log.Println("Error decoding recipe creation request:", err)
		return
	}

	ingredients := make([]recipes.RecipeIngredient, 0, len(dto.Ingredients))
	for _, ingredientDto := range dto.Ingredients {
		ingredients = append(ingredients, recipes.RecipeIngredient{
			Ingredient: ingredientDto.Ingredient,
			Amount:     ingredientDto.Amount,
			Unit:       ingredientDto.Unit,
		})
	}

	instructions := make([]recipes.RecipeStep, 0, len(dto.Instructions))
	for _, stepDto := range dto.Instructions {
		instructions = append(instructions, recipes.RecipeStep{
			Text: stepDto.Text,
			Time: stepDto.Time,
		})
	}

	newRecipe := &recipes.Recipe{
		Title:        dto.Title,
		Description:  dto.Description,
		Ingredients:  ingredients,
		Instructions: instructions,
	}

	if err := h.service.CreateRecipe(dto.Username, newRecipe); err != nil {
		shared.WriteError(w, err)
		log.Println("Error creating recipe:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *RecipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	title := q.Get("title")

	var foundRecipes []recipes.Recipe
	var err error

	if title != "" {

		var recipe *recipes.Recipe
		recipe, err = h.service.GetRecipeWithTitle(title)
		foundRecipes = append(foundRecipes, *recipe)
	} else {
		foundRecipes = h.service.GetAllRecipes()
	}

	if err != nil {
		shared.WriteError(w, err)
		return
	}

	recipeDtos := make([]RecipeDto, 0, len(foundRecipes))

	for _, recipe := range foundRecipes {
		user, err := h.userService.GetUserByUserId(recipe.UserID)
		if err != nil {
			shared.WriteError(w, err)
			return
		}

		ingredientsDtos := make([]RecipeIngredientDto, 0, len(recipe.Ingredients))
		for _, ingredient := range recipe.Ingredients {
			ingredientsDtos = append(ingredientsDtos, RecipeIngredientDto{
				Ingredient: ingredient.Ingredient,
				Amount:     ingredient.Amount,
				Unit:       ingredient.Unit,
			})
		}

		instructionsDtos := make([]RecipeStepDto, 0, len(recipe.Instructions))
		for _, step := range recipe.Instructions {
			instructionsDtos = append(instructionsDtos, RecipeStepDto{
				Text: step.Text,
				Time: step.Time,
			})
		}

		recipeDtos = append(recipeDtos, RecipeDto{
			Title:        recipe.Title,
			Username:     user.Username,
			Description:  recipe.Description,
			Ingredients:  ingredientsDtos,
			Instructions: instructionsDtos,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(recipeDtos); err != nil {
		shared.WriteError(w, err)
		return
	}
}
