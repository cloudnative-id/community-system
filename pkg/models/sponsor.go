package models

import (
	"github.com/google/uuid"
)

type Sponsor struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Image string    `json:"image"`
}
