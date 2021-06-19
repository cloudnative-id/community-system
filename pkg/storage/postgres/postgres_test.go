package postgres_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/cloudnative-id/community-system/pkg/storage/postgres"

	postgres_mock "github.com/cloudnative-id/community-system/pkg/storage/postgres/mock"
)

func TestPostgresMigrateMeetup(t *testing.T) {
	assert := assert.New(t)
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockPostgresClient := postgres_mock.NewMockPostgresClientInterface(controller)
	client := postgres.PostgresClient{
		Client: mockPostgresClient,
	}

	mockPostgresClient.EXPECT().AutoMigrate(&models.Meetup{}).Return(nil)

	err := client.MigrateMeetup()
	assert.Nil(err)
}

func TestPostgresMigrateSpeaker(t *testing.T) {
	assert := assert.New(t)
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockPostgresClient := postgres_mock.NewMockPostgresClientInterface(controller)
	client := postgres.PostgresClient{
		Client: mockPostgresClient,
	}

	mockPostgresClient.EXPECT().AutoMigrate(&models.Speaker{}).Return(nil)

	err := client.MigrateSpeaker()
	assert.Nil(err)
}

func TestPostgresMigrateSponsor(t *testing.T) {
	assert := assert.New(t)
	controller := gomock.NewController(t)
	defer controller.Finish()

	mockPostgresClient := postgres_mock.NewMockPostgresClientInterface(controller)
	client := postgres.PostgresClient{
		Client: mockPostgresClient,
	}

	mockPostgresClient.EXPECT().AutoMigrate(&models.Sponsor{}).Return(nil)

	err := client.MigrateSponsor()
	assert.Nil(err)
}
