package models

import (
	"github.com/google/uuid"
)

type Speaker struct {
	UUID     uuid.UUID `json:"uuid"`
	Name     string    `json:"name"`
	Position string    `json:"position"`
	Company  string    `json:"company"`
	Title    string    `json:"title"`
	Image    string    `json:"image"`
}
