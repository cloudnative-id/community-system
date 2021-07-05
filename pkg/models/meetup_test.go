package models_test

import (
	"testing"
	"time"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMeetupSanitize(t *testing.T) {
	t.Run("give empty, return error", func(t *testing.T) {
		assert := assert.New(t)
		meetup := models.Meetup{}

		assert.Error(meetup.Sanitize())
	})

	t.Run("missing uuid, return error", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		city := "bar"
		location := "foo"
		duration := int64(90)
		time := strfmt.DateTime(time.Now())
		tags := []string{"foo", "bar"}
		meetup := models.Meetup{
			Country:  &country,
			City:     &city,
			Location: &location,
			Duration: &duration,
			Time:     &time,
			Tags:     tags,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("missing country, return error", func(t *testing.T) {
		assert := assert.New(t)
		city := "bar"
		location := "foo"
		duration := int64(90)
		time := strfmt.DateTime(time.Now())
		tags := []string{"foo", "bar"}
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			City:     &city,
			Location: &location,
			Duration: &duration,
			Time:     &time,
			Tags:     tags,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("missing city, return error", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		location := "foo"
		duration := int64(90)
		time := strfmt.DateTime(time.Now())
		tags := []string{"foo", "bar"}
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			Country:  &country,
			Location: &location,
			Duration: &duration,
			Time:     &time,
			Tags:     tags,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("missing location, return error", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		city := "bar"
		duration := int64(90)
		time := strfmt.DateTime(time.Now())
		tags := []string{"foo", "bar"}
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			Country:  &country,
			City:     &city,
			Duration: &duration,
			Time:     &time,
			Tags:     tags,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("missing duration, return error", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		city := "bar"
		location := "foo"
		time := strfmt.DateTime(time.Now())
		tags := []string{"foo", "bar"}
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			Country:  &country,
			City:     &city,
			Location: &location,
			Time:     &time,
			Tags:     tags,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("missing time, return error", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		city := "bar"
		location := "foo"
		duration := int64(90)
		tags := []string{"foo", "bar"}
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			Country:  &country,
			City:     &city,
			Location: &location,
			Duration: &duration,
			Tags:     tags,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("missing tags, return error", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		city := "bar"
		location := "foo"
		duration := int64(90)
		time := strfmt.DateTime(time.Now())
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			Country:  &country,
			City:     &city,
			Location: &location,
			Duration: &duration,
			Time:     &time,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("duration is minus, return error", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		city := "bar"
		location := "foo"
		duration := int64(-90)
		time := strfmt.DateTime(time.Now())
		tags := []string{"foo", "bar"}
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			Country:  &country,
			City:     &city,
			Location: &location,
			Duration: &duration,
			Time:     &time,
			Tags:     tags,
		}

		assert.Error(meetup.Sanitize())
	})

	t.Run("give correct value, return nil", func(t *testing.T) {
		assert := assert.New(t)
		country := "foo"
		city := "bar"
		location := "foo"
		duration := int64(90)
		time := strfmt.DateTime(time.Now())
		tags := []string{"foo", "bar"}
		meetupUUID := uuid.New()
		meetup := models.Meetup{
			UUID:     &meetupUUID,
			Country:  &country,
			City:     &city,
			Location: &location,
			Duration: &duration,
			Time:     &time,
			Tags:     tags,
		}

		assert.Nil(meetup.Sanitize())
	})
}

func TestMeetupToAPIMeetup(t *testing.T) {
	assert := assert.New(t)
	country := "foo"
	city := "bar"
	location := "foo"
	duration := int64(90)
	time := strfmt.DateTime(time.Now())
	tags := []string{"foo", "bar"}
	meetupUUID := uuid.New()
	meetup := models.Meetup{
		UUID:     &meetupUUID,
		Country:  &country,
		City:     &city,
		Location: &location,
		Duration: &duration,
		Time:     &time,
		Tags:     tags,
	}

	meetupAPIModel := &api_model.Meetup{
		UUID:     meetupUUID.String(),
		Country:  &country,
		City:     &city,
		Location: &location,
		Duration: &duration,
		Time:     &time,
		Tags:     tags,
	}
	assert.Equal(meetupAPIModel, meetup.ToAPIMeetup())
}
