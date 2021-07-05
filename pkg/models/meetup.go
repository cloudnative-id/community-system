package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	api_model "github.com/cloudnative-id/community-system/gen/models"
)

type Meetup struct {
	Model
	UUID            *uuid.UUID       `gorm:"primaryKey;type:uuid;not null" json:"uuid"`
	Country         *string          `gorm:"type:varchar(255);not null" json:"country"`
	City            *string          `gorm:"type:varchar(255);not null" json:"city"`
	Location        *string          `gorm:"type:varchar(255);not null" json:"location"`
	Duration        *int64           `gorm:"not null" json:"duration"`
	Time            *strfmt.DateTime `gorm:"not null" json:"time"`
	Tags            Tags             `gorm:"type:jsonb;not null;" json:"tags"`
	RegistrationURL string           `gorm:"type:varchar(255);" json:"registrationUrl"`
	Image           string           `gorm:"type:varchar(255);" json:"image"`
	Speakers        UUIDs            `gorm:"type:jsonb;" json:"speaker"`
	Sponsors        UUIDs            `gorm:"type:jsonb;" json:"sponsors"`
	Status          bool             `json:"status"`
}

func (m Meetup) sanitize() error {
	if m.UUID == nil {
		return fmt.Errorf("UUID cannot empty")
	}

	if m.Country == nil {
		return fmt.Errorf("Country cannot empty")
	}

	if m.City == nil {
		return fmt.Errorf("City cannot empty")
	}

	if m.Location == nil {
		return fmt.Errorf("Location cannot empty")
	}

	if m.Duration == nil {
		return fmt.Errorf("Duration cannot empty")
	}

	if *m.Duration < 0 {
		return fmt.Errorf("Duration cannot be 0 or less than 0")
	}

	if m.Time == nil {
		return fmt.Errorf("Location cannot empty")
	}

	if len(m.Tags) == 0 {
		return fmt.Errorf("Tags cannot empty")
	}

	return nil
}

func (m Meetup) ToAPIMeetup() *api_model.Meetup {
	return &api_model.Meetup{
		UUID:            m.UUID.String(),
		Country:         m.Country,
		City:            m.City,
		Location:        m.Location,
		Duration:        m.Duration,
		Time:            m.Time,
		Tags:            m.Tags,
		RegistrationURL: m.RegistrationURL,
		Image:           m.Image,
		Speakers:        m.Speakers.ArrayString(),
		Sponsors:        m.Sponsors.ArrayString(),
		Status:          m.Status,
	}
}

type Tags []string

// implements sql.Scanner interface
func (t *Tags) Scan(value interface{}) error {
	byte, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(byte, &t)
}

// implement driver.Valuer interface
func (t Tags) Value() (driver.Value, error) {
	return json.Marshal(t)
}

type UUIDs []uuid.UUID

// implements sql.Scanner interface
func (u *UUIDs) Scan(value interface{}) error {
	byte, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(byte, &u)
}

// implement driver.Valuer interface
func (u UUIDs) Value() (driver.Value, error) {
	return json.Marshal(u)
}

func (u UUIDs) ArrayString() []string {
	var output []string
	for _, uuid := range u {
		output = append(output, uuid.String())
	}

	return output
}

func NewUUIDs(uuidStringList []string) (UUIDs, error) {
	var uuids UUIDs
	if len(uuidStringList) > 0 {
		for _, uuidString := range uuidStringList {
			uuid, err := uuid.Parse(uuidString)
			if err != nil {
				return uuids, err
			}

			uuids = append(uuids, uuid)
		}
	}

	return uuids, nil
}
