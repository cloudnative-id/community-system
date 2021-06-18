package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Meetup struct {
	Model
	UUID            uuid.UUID       `gorm:"primaryKey;type:uuid" json:"uuid"`
	Country         string          `gorm:"type:varchar(255);not null" json:"country"`
	City            string          `gorm:"type:varchar(255);not null" json:"city"`
	Location        string          `gorm:"type:varchar(255);not null" json:"location"`
	Year            int64           `gorm:"type:varchar(255);not null" json:"year"`
	Month           int64           `gorm:"type:varchar(255);not null" json:"month"`
	Day             int64           `gorm:"type:varchar(255);not null" json:"day"`
	Time            strfmt.DateTime `gorm:"not null" json:"time"`
	RegistrationURL string          `gorm:"type:varchar(255);" json:"registrationUrl"`
	Image           string          `gorm:"type:varchar(255);" json:"image"`
	Tags            Tags            `gorm:"type:jsonb;" json:"tags"`
	Speaker         UUIDs           `gorm:"type:jsonb;" json:"speaker"`
	Sponsors        UUIDs           `gorm:"type:jsonb;" json:"sponsors"`
	Status          bool            `json:"status"`
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
