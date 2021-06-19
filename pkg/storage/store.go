package storage

import "github.com/cloudnative-id/community-system/pkg/models"

type Storage interface {
	MigrateMeetup() error
	MigrateSpeaker() error
	MigrateSponsor() error
	WriteMeetup(meetup models.Meetup) error
	GetMeetup(uuid string) (models.Meetup, bool, error)
	GetMeetups() ([]models.Meetup, bool, error)
	DeleteMeetup(meetup models.Meetup) error
}
