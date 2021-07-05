package models

import (
	"fmt"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	"github.com/google/uuid"
)

type Speaker struct {
	Model
	UUID       *uuid.UUID `gorm:"primaryKey;type:uuid" json:"uuid"`
	MeetupUUID *uuid.UUID `gorm:"type:uuid;not null" json:"meetup_uuid"`
	Name       *string    `gorm:"type:varchar(255);not null" json:"name"`
	Position   *string    `gorm:"type:varchar(255);not null" json:"position"`
	Company    *string    `gorm:"type:varchar(255);not null" json:"company"`
	Title      *string    `gorm:"type:varchar(255);not null" json:"title"`
	Image      *string    `gorm:"type:varchar(255);not null" json:"image"`
}

func (s Speaker) Sanitize() error {
	if s.UUID == nil {
		return fmt.Errorf("uuid cannot empty")
	}

	if s.MeetupUUID == nil {
		return fmt.Errorf("meetup uuid cannot empty")
	}

	if s.Name == nil {
		return fmt.Errorf("name cannot empty")
	}

	if s.Position == nil {
		return fmt.Errorf("position cannot empty")
	}

	if s.Company == nil {
		return fmt.Errorf("company cannot empty")
	}

	if s.Title == nil {
		return fmt.Errorf("position cannot empty")
	}

	if s.Image == nil {
		return fmt.Errorf("image cannot empty")
	}

	return nil
}

func (s Speaker) ToAPIMeetup() *api_model.Speaker {
	return &api_model.Speaker{
		UUID:     s.UUID.String(),
		Name:     s.Name,
		Position: s.Position,
		Company:  s.Company,
		Title:    s.Title,
		Image:    s.Image,
	}
}
