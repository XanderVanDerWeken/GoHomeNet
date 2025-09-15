package chores

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/xandervanderweken/GoHomeNet/internal/shared"
	"github.com/xandervanderweken/GoHomeNet/internal/users"
)

type ChoreHandler struct {
	service     Service
	userService users.Service
}

func NewChoreHandler(service Service, userService users.Service) *ChoreHandler {
	return &ChoreHandler{service: service, userService: userService}
}

func (h *ChoreHandler) PostNewChore(w http.ResponseWriter, r *http.Request) {
	var dto NewChoreDto
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		shared.WriteError(w, err)
		return
	}

	if err := h.service.CreateChore(dto.Username, dto.Title, dto.Notes, &dto.DueDate); err != nil {
		shared.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ChoreHandler) GetChores(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	username := q.Get("username")

	var chores []Chore
	var err error

	if username != "" {
		chores, err = h.service.GetChoresByUsername(username)
	} else {
		chores = h.service.GetAllChores()
	}
	if err != nil {
		shared.WriteError(w, err)
		return
	}

	choreDtos := make([]ChoreDto, 0, len(chores))

	for _, chore := range chores {
		uName := username
		if uName == "" {
			user, err := h.userService.GetUserByUserId(chore.UserID)
			if err != nil {
				shared.WriteError(w, err)
				return
			}
			uName = user.Username
		}

		choreDtos = append(choreDtos, ChoreDto{
			Username:  uName,
			Title:     chore.Title,
			Notes:     chore.Notes,
			DueDate:   chore.DueDate,
			Completed: chore.Completed,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(choreDtos); err != nil {
		shared.WriteError(w, err)
		return
	}
}

func (h *ChoreHandler) PutChoreComplete(w http.ResponseWriter, r *http.Request) {
	choreId := chi.URLParam(r, "choreId")

	u64, err := strconv.ParseUint(choreId, 10, 32)
	if err != nil {
		shared.WriteError(w, fmt.Errorf("invalid chore id: %w", err))
		return
	}

	err = h.service.CompleteChore(uint(u64))
	if err != nil {
		shared.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChoreHandler) DeleteChore(w http.ResponseWriter, r *http.Request) {
	choreId := chi.URLParam(r, "choreId")

	u64, err := strconv.ParseUint(choreId, 10, 32)
	if err != nil {
		shared.WriteError(w, fmt.Errorf("invalid chore id: %w", err))
		return
	}

	err = h.service.DeleteChore(uint(u64))
	if err != nil {
		shared.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
