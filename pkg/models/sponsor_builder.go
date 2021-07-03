package models

import (
	"github.com/google/uuid"
)

type SponsorBuilder struct {
	sponsor Sponsor
}

func NewSponsorBuilder() *SponsorBuilder {
	return &SponsorBuilder{
		sponsor: Sponsor{},
	}
}

func (m *SponsorBuilder) SetUUID(uuid *uuid.UUID) *SponsorBuilder {
	m.sponsor.UUID = uuid
	return m
}

func (m *SponsorBuilder) SetName(name *string) *SponsorBuilder {
	m.sponsor.Name = name
	return m
}

func (m *SponsorBuilder) SetImage(image *string) *SponsorBuilder {
	m.sponsor.Image = image
	return m
}

func (m *SponsorBuilder) Build() (Sponsor, error) {
	err := m.sponsor.sanitize()
	if err != nil {
		return Sponsor{}, err
	}

	return m.sponsor, nil
}
