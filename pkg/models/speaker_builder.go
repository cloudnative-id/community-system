package models

import (
	"github.com/google/uuid"
)

type SpeakerBuilder struct {
	speaker Speaker
}

func NewSpeakerBuilder() *SpeakerBuilder {
	return &SpeakerBuilder{
		speaker: Speaker{},
	}
}

func (m *SpeakerBuilder) SetUUID(uuid *uuid.UUID) *SpeakerBuilder {
	m.speaker.UUID = uuid
	return m
}

func (m *SpeakerBuilder) SetMeetupUUID(uuid *uuid.UUID) *SpeakerBuilder {
	m.speaker.MeetupUUID = uuid
	return m
}

func (m *SpeakerBuilder) SetName(name *string) *SpeakerBuilder {
	m.speaker.Name = name
	return m
}

func (m *SpeakerBuilder) SetPosition(position *string) *SpeakerBuilder {
	m.speaker.Position = position
	return m
}

func (m *SpeakerBuilder) SetCompany(company *string) *SpeakerBuilder {
	m.speaker.Company = company
	return m
}

func (m *SpeakerBuilder) SetTitle(title *string) *SpeakerBuilder {
	m.speaker.Title = title
	return m
}

func (m *SpeakerBuilder) SetImage(image *string) *SpeakerBuilder {
	m.speaker.Image = image
	return m
}

func (m *SpeakerBuilder) Build() (Speaker, error) {
	err := m.speaker.Sanitize()
	if err != nil {
		return Speaker{}, err
	}

	return m.speaker, nil
}
