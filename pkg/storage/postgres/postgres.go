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

func (p PostgresClient) GetMeetup(uuid string) (models.Meetup, bool, error) {
	var meetup models.Meetup
	tx := p.Client.Where("uuid = ?", uuid).First(&meetup)

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
