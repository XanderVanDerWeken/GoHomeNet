package recipes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type RecipeHandler struct {
	service     Service
	userService users.Service
}

func NewRecipeHandler(service Service, userSerivce users.Service) *RecipeHandler {
	return &RecipeHandler{
		service:     service,
		userService: userSerivce,
	}
}

func (h *RecipeHandler) PostNewRecipe(w http.ResponseWriter, r *http.Request) {
	var dto struct {
		Title        string                `json:"title"`
		Username     string                `json:"author"`
		Description  string                `json:"description"`
		Ingredients  []RecipeIngredientDto `json:"ingredients"`
		Instructions []RecipeStepDto       `json:"instructions"`
	}
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		writeError(w, err)
		log.Println("Error decoding recipe creation request:", err)
		return
	}

	if err := h.service.CreateRecipe(dto.Username, dto.Title); err != nil {
		writeError(w, err)
		log.Println("Error creating recipe:", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *RecipeHandler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	title := q.Get("title")

	var recipes []Recipe
	var err error

	if title != "" {
		var recipe *Recipe
		recipe, err = h.service.GetRecipeWithTitle(title)
		recipes = append(recipes, *recipe)
	} else {
		recipes = h.service.GetAllRecipes()
	}

	if err != nil {
		writeError(w, err)
		return
	}

	recipeDtos := make([]RecipeDto, 0, len(recipes))

	for _, recipe := range recipes {
		user, err := h.userService.GetUserByUserId(recipe.UserID)
		if err != nil {
			writeError(w, err)
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
		writeError(w, err)
		return
	}
}

func writeError(w http.ResponseWriter, err error) {
	if appErr, ok := err.(*shared.AppError); ok {
		http.Error(w, appErr.Message, appErr.Status)
		return
	}
	http.Error(w, shared.ErrInternal.Message, shared.ErrInternal.Status)
}
