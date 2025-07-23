package chores

import "time"

type ChoreDto struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	DueDate     time.Time `json:"dueDate"`
	IsDone      bool      `json:"isDone"`
}

type CreateChoreRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description *string   `json:"description,omitempty"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
}
