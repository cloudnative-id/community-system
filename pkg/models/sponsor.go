package models

import (
	"fmt"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	"github.com/google/uuid"
)

type Sponsor struct {
	Model
	UUID  *uuid.UUID `gorm:"primaryKey;type:uuid" json:"uuid"`
	Name  *string    `gorm:"type:varchar(255);not null" json:"name"`
	Image *string    `gorm:"type:varchar(255);not null" json:"image"`
}

func (s Sponsor) Sanitize() error {
	if s.UUID == nil {
		return fmt.Errorf("uuid cannot empty")
	}

	if s.Name == nil {
		return fmt.Errorf("name cannot empty")
	}

	if s.Image == nil {
		return fmt.Errorf("image cannot empty")
	}

	return nil
}

func (s Sponsor) ToAPIMeetup() *api_model.Sponsor {
	return &api_model.Sponsor{
		UUID:  s.UUID.String(),
		Name:  s.Name,
		Image: s.Image,
	}
}
