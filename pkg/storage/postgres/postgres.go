package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/cloudnative-id/community-system/pkg/common"
	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/cloudnative-id/community-system/pkg/settings"
	"github.com/cloudnative-id/community-system/pkg/storage"
)

type PostgresClientInterface interface {
	Clauses(conds ...clause.Expression) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	AutoMigrate(dst ...interface{}) error
}

type PostgresClient struct {
	Client PostgresClientInterface
}

func (p PostgresClient) MigrateMeetup() error {
	return p.Client.AutoMigrate(&models.Meetup{})
}

func (p PostgresClient) MigrateSpeaker() error {
	return p.Client.AutoMigrate(&models.Speaker{})
}

func (p PostgresClient) MigrateSponsor() error {
	return p.Client.AutoMigrate(&models.Sponsor{})
}

func (p PostgresClient) WriteMeetup(meetup models.Meetup) error {
	tx := p.Client.Clauses(clause.OnConflict{UpdateAll: true}).Create(&meetup)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p PostgresClient) GetMeetup(meetupUUID string) (models.Meetup, bool, error) {
	var meetup models.Meetup
	tx := p.Client.Where("uuid = ?", meetupUUID).First(&meetup)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return models.Meetup{}, false, nil
		}

		return models.Meetup{}, false, tx.Error
	}

	return meetup, true, nil
}

func (p PostgresClient) GetMeetups() ([]models.Meetup, bool, error) {
	var meetups []models.Meetup
	tx := p.Client.Find(&meetups)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return []models.Meetup{}, false, nil
		}

		return []models.Meetup{}, false, tx.Error
	}

	return meetups, true, nil
}

func (p PostgresClient) DeleteMeetup(meetup models.Meetup) error {
	return common.NotImplemented
}

func (p PostgresClient) WriteSpeaker(speaker models.Speaker) error {
	tx := p.Client.Clauses(clause.OnConflict{UpdateAll: true}).Create(&speaker)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p PostgresClient) GetSpeaker(meetupUUID string, speakerUUID string) (models.Speaker, bool, error) {
	var speaker models.Speaker
	tx := p.Client.Where("meetup_uuid = ? AND uuid = ?", meetupUUID, speakerUUID).First(&speaker)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return models.Speaker{}, false, nil
		}

		return models.Speaker{}, false, tx.Error
	}

	return speaker, true, nil
}

func (p PostgresClient) GetSpeakers(meetupUUID string) ([]models.Speaker, bool, error) {
	var speakers []models.Speaker
	tx := p.Client.Where("meetup_uuid = ?", meetupUUID).Find(&speakers)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return []models.Speaker{}, false, nil
		}

		return []models.Speaker{}, false, tx.Error
	}

	return speakers, true, nil
}

func (p PostgresClient) DeleteSpeaker(speaker models.Speaker) error {
	return common.NotImplemented
}

func (p PostgresClient) WriteSponsor(sponsor models.Sponsor) error {
	tx := p.Client.Clauses(clause.OnConflict{UpdateAll: true}).Create(&sponsor)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p PostgresClient) GetSponsor(sponsorUUID string) (models.Sponsor, bool, error) {
	var sponsor models.Sponsor
	tx := p.Client.Where("uuid = ?", sponsorUUID).First(&sponsor)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return models.Sponsor{}, false, nil
		}

		return models.Sponsor{}, false, tx.Error
	}

	return sponsor, true, nil
}

func (p PostgresClient) GetSponsors() ([]models.Sponsor, bool, error) {
	var sponsors []models.Sponsor
	tx := p.Client.Find(&sponsors)

	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return []models.Sponsor{}, false, nil
		}

		return []models.Sponsor{}, false, tx.Error
	}

	return sponsors, true, nil
}

func (p PostgresClient) DeleteSponsor(sponsor models.Sponsor) error {
	return common.NotImplemented
}

func NewPostgres(settings settings.Settings) storage.Storage {
	if settings.DatabasePort == "" {
		settings.DatabasePort = "5432"
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s DB.name=%s password=%s sslmode=disable",
		settings.DatabaseHost,
		settings.DatabasePort,
		settings.DatabaseUser,
		settings.DatabaseName,
		settings.DatabasePassword)
	client, _ := gorm.Open(postgres.New(postgres.Config{DSN: dsn}), &gorm.Config{})

	return PostgresClient{
		Client: client,
	}
}
