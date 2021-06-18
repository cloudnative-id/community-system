package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type Meetup struct {
	UUID            uuid.UUID       `json:"uuid"`
	Country         string          `json:"country"`
	City            string          `json:"city"`
	Year            int64           `json:"year"`
	Month           int64           `json:"month"`
	Day             int64           `json:"day"`
	Time            strfmt.DateTime `json:"time"`
	Location        string          `json:"location"`
	RegistrationURL string          `json:"registrationUrl"`
	Image           string          `json:"image"`
	Tags            []string        `json:"tags"`
	Speaker         []Speaker       `json:"speaker"`
	Sponsors        []Sponsor       `json:"sponsors"`
	Status          bool            `json:"status"`
}
