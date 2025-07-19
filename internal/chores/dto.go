package chores

import "time"

type CreateChoreRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description *string   `json:"description,omitempty"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
}
