package finances

import (
	"net/http"

	"github.com/xandervanderweken/GoHomeNet/internal/shared"
)

var (
	ErrCategoryAlreadyExists = shared.NewAppError("CATEGORY_EXISTS", "Category with name already exists", http.StatusConflict)
)
