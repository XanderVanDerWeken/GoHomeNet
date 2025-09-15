package recipes

import (
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

var (
	ErrRecipeAlreadyExists = shared.NewAppError("RECIPE_EXISTS", "Recipe with title already exists", http.StatusConflict)
)
