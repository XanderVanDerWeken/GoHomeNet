package cards

import "time"

type NewCardDto struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type CardDto struct {
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	ExpiresAt time.Time `json:"expiresAt"`
}
