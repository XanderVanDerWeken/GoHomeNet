package cards

import "time"

type CreateCardRequest struct {
	Name    string    `json:"name"`
	DueDate time.Time `json:"due_date"`
}

type CardDto struct {
	ID      uint      `json:"id"`
	Name    string    `json:"name"`
	DueDate time.Time `json:"due_date"`
}
