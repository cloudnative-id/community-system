package models_test

import (
	"testing"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	api_model "github.com/cloudnative-id/community-system/gen/models"
)

func TestSpeakerSanitize(t *testing.T) {
	t.Run("give empty, return error", func(t *testing.T) {
		assert := assert.New(t)
		speaker := models.Speaker{}

		assert.Error(speaker.Sanitize())
	})

	t.Run("missing speaker UUID, return error", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		position := "bar"
		company := "foo"
		title := "bar"
		image := "foo"
		meeetupUUID := uuid.New()
		speaker := models.Speaker{
			MeetupUUID: &meeetupUUID,
			Name:       &name,
			Position:   &position,
			Company:    &company,
			Title:      &title,
			Image:      &image,
		}

		assert.Error(speaker.Sanitize())
	})

	t.Run("missing meetup UUID, return error", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		position := "bar"
		company := "foo"
		title := "bar"
		image := "foo"
		speakerUUID := uuid.New()
		speaker := models.Speaker{
			UUID:     &speakerUUID,
			Name:     &name,
			Position: &position,
			Company:  &company,
			Title:    &title,
			Image:    &image,
		}

		assert.Error(speaker.Sanitize())
	})

	t.Run("missing name, return error", func(t *testing.T) {
		assert := assert.New(t)
		position := "bar"
		company := "foo"
		title := "bar"
		image := "foo"
		speakerUUID := uuid.New()
		meeetupUUID := uuid.New()
		speaker := models.Speaker{
			UUID:       &speakerUUID,
			MeetupUUID: &meeetupUUID,
			Position:   &position,
			Company:    &company,
			Title:      &title,
			Image:      &image,
		}

		assert.Error(speaker.Sanitize())
	})

	t.Run("missing company, return error", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		position := "bar"
		title := "bar"
		image := "foo"
		speakerUUID := uuid.New()
		meeetupUUID := uuid.New()
		speaker := models.Speaker{
			UUID:       &speakerUUID,
			MeetupUUID: &meeetupUUID,
			Name:       &name,
			Position:   &position,
			Title:      &title,
			Image:      &image,
		}

		assert.Error(speaker.Sanitize())
	})

	t.Run("missing position, return error", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		company := "foo"
		title := "bar"
		image := "foo"
		speakerUUID := uuid.New()
		meeetupUUID := uuid.New()
		speaker := models.Speaker{
			UUID:       &speakerUUID,
			MeetupUUID: &meeetupUUID,
			Name:       &name,
			Company:    &company,
			Title:      &title,
			Image:      &image,
		}

		assert.Error(speaker.Sanitize())
	})

	t.Run("missing title, return error", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		position := "bar"
		company := "foo"
		image := "foo"
		speakerUUID := uuid.New()
		meeetupUUID := uuid.New()
		speaker := models.Speaker{
			UUID:       &speakerUUID,
			MeetupUUID: &meeetupUUID,
			Name:       &name,
			Position:   &position,
			Company:    &company,
			Image:      &image,
		}

		assert.Error(speaker.Sanitize())
	})

	t.Run("missing image, return error", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		position := "bar"
		company := "foo"
		title := "bar"
		speakerUUID := uuid.New()
		meeetupUUID := uuid.New()
		speaker := models.Speaker{
			UUID:       &speakerUUID,
			MeetupUUID: &meeetupUUID,
			Name:       &name,
			Position:   &position,
			Company:    &company,
			Title:      &title,
		}

		assert.Error(speaker.Sanitize())
	})

	t.Run("give correct value, return nil", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		position := "bar"
		company := "foo"
		title := "bar"
		image := "foo"
		speakerUUID := uuid.New()
		meeetupUUID := uuid.New()
		speaker := models.Speaker{
			UUID:       &speakerUUID,
			MeetupUUID: &meeetupUUID,
			Name:       &name,
			Position:   &position,
			Company:    &company,
			Title:      &title,
			Image:      &image,
		}

		assert.Nil(speaker.Sanitize())
	})
}

func TestSpeakerToAPIMeetup(t *testing.T) {
	assert := assert.New(t)
	name := "foo"
	position := "bar"
	company := "foo"
	title := "bar"
	image := "foo"
	speakerUUID := uuid.New()
	meeetupUUID := uuid.New()
	speaker := models.Speaker{
		UUID:       &speakerUUID,
		MeetupUUID: &meeetupUUID,
		Name:       &name,
		Position:   &position,
		Company:    &company,
		Title:      &title,
		Image:      &image,
	}

	speakerAPIModel := &api_model.Speaker{
		UUID:     speakerUUID.String(),
		Name:     &name,
		Image:    &image,
		Position: &position,
		Company:  &company,
		Title:    &title,
	}
	assert.Equal(speakerAPIModel, speaker.ToAPIMeetup())
}
