package chores

import "time"

type NewChoreDto struct {
	Username string    `json:"username"`
	Title    string    `json:"title"`
	Notes    string    `json:"notes"`
	DueDate  time.Time `json:"dueDate"`
}

type ChoreDto struct {
	Username  string     `json:"username"`
	Title     string     `json:"title"`
	Notes     string     `json:"notes"`
	DueDate   *time.Time `json:"dueDate,omitempty"`
	Completed bool       `json:"completed"`
}
