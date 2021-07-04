package storage

import "github.com/cloudnative-id/community-system/pkg/models"

type Storage interface {
	MigrateMeetup() error
	MigrateSpeaker() error
	MigrateSponsor() error

	WriteMeetup(meetup models.Meetup) error
	GetMeetup(meetupUUID string) (models.Meetup, bool, error)
	GetMeetups() ([]models.Meetup, bool, error)
	DeleteMeetup(meetup models.Meetup) error

	WriteSpeaker(speaker models.Speaker) error
	GetSpeakers(meetupUUID string) ([]models.Speaker, bool, error)
	DeleteSpeaker(speaker models.Speaker) error

	WriteSponsor(sponsor models.Sponsor) error
	GetSponsor(sponsorUUID string) (models.Sponsor, bool, error)
	GetSponsors() ([]models.Sponsor, bool, error)
	DeleteSponsor(sponsor models.Sponsor) error
}
