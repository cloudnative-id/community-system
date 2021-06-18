package storage

import "github.com/cloudnative-id/community-system/pkg/models"

type Storage interface {
	WriteMeetup(meetup models.Meetup) error
	GetMeetup() (models.Meetup, error)
	GetMeetups() ([]models.Meetup, error)
	DeleteMeetup(meetup models.Meetup) error
}
