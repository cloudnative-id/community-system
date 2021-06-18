package models

import (
	"github.com/google/uuid"
)

type Sponsor struct {
	Model
	UUID  uuid.UUID `gorm:"primaryKey;type:uuid" json:"uuid"`
	Name  string    `gorm:"type:varchar(255);not null" json:"name"`
	Image string    `gorm:"type:varchar(255);not null" json:"image"`
}
