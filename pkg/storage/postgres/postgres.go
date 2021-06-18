package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/cloudnative-id/community-system/pkg/settings"
)

type PostgresClientInterface interface {
	Clauses(conds ...clause.Expression) (tx *gorm.DB)
	Create(value interface{}) (tx *gorm.DB)
	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
	Save(value interface{}) (tx *gorm.DB)
	AutoMigrate(dst ...interface{}) error
}

type PostgresClient struct {
	Client PostgresClientInterface
}

func (p PostgresClient) WriteMeetup(meetup models.Meetup) error {
	tx := p.Client.Clauses(clause.OnConflict{UpdateAll: true}).Create(&meetup)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (p PostgresClient) GetMeetup() (models.Meetup, error) {
	return models.Meetup{}, nil
}

func (p PostgresClient) GetMeetups() ([]models.Meetup, error) {
	return []models.Meetup{}, nil
}

func (p PostgresClient) DeleteMeetup(meetup models.Meetup) error {
	return nil
}

func NewPostgres(settings settings.Settings) PostgresClient {
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
