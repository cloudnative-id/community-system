package models

import (
	"github.com/google/uuid"
)

type Speaker struct {
	Model
	UUID     uuid.UUID `gorm:"primaryKey;type:uuid" json:"uuid"`
	Name     string    `gorm:"type:varchar(255);not null" json:"name"`
	Position string    `gorm:"type:varchar(255);not null" json:"position"`
	Company  string    `gorm:"type:varchar(255);not null" json:"company"`
	Title    string    `gorm:"type:varchar(255);not null" json:"title"`
	Image    string    `gorm:"type:varchar(255);not null" json:"image"`
}
