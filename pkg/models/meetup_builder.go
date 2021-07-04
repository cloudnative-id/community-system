package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type MeetupBuilder struct {
	meetup Meetup
}

func NewMeetupBuilder() *MeetupBuilder {
	return &MeetupBuilder{
		meetup: Meetup{},
	}
}

func (m *MeetupBuilder) SetUUID(uuid *uuid.UUID) *MeetupBuilder {
	m.meetup.UUID = uuid
	return m
}

func (m *MeetupBuilder) SetCountry(country *string) *MeetupBuilder {
	m.meetup.Country = country
	return m
}

func (m *MeetupBuilder) SetCity(city *string) *MeetupBuilder {
	m.meetup.City = city
	return m
}

func (m *MeetupBuilder) SetLocation(location *string) *MeetupBuilder {
	m.meetup.Location = location
	return m
}

func (m *MeetupBuilder) SetTags(tags Tags) *MeetupBuilder {
	m.meetup.Tags = tags
	return m
}

func (m *MeetupBuilder) SetDuration(duration *int64) *MeetupBuilder {
	m.meetup.Duration = duration
	return m
}

func (m *MeetupBuilder) SetTime(time *strfmt.DateTime) *MeetupBuilder {
	m.meetup.Time = time
	return m
}

func (m *MeetupBuilder) SetRegistrationURL(registrationURL string) *MeetupBuilder {
	m.meetup.RegistrationURL = registrationURL
	return m
}

func (m *MeetupBuilder) SetImage(image string) *MeetupBuilder {
	m.meetup.Image = image
	return m
}

func (m *MeetupBuilder) SetSpeakers(speakers UUIDs) *MeetupBuilder {
	m.meetup.Speakers = speakers
	return m
}

func (m *MeetupBuilder) SetSponsors(sponsors UUIDs) *MeetupBuilder {
	m.meetup.Sponsors = sponsors
	return m
}

func (m *MeetupBuilder) SetStatus(status bool) *MeetupBuilder {
	m.meetup.Status = status
	return m
}

func (m *MeetupBuilder) Build() (Meetup, error) {
	err := m.meetup.sanitize()
	if err != nil {
		return Meetup{}, err
	}

	return m.meetup, nil
}
