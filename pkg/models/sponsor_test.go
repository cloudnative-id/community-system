package models_test

import (
	"testing"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	api_model "github.com/cloudnative-id/community-system/gen/models"
)

func TestSponsorSanitize(t *testing.T) {
	t.Run("give empty, return error", func(t *testing.T) {
		assert := assert.New(t)
		sponsor := models.Sponsor{}

		assert.Error(sponsor.Sanitize())
	})

	t.Run("missing uuid, return error", func(t *testing.T) {
		assert := assert.New(t)
		name := "foo"
		image := "bar"
		sponsor := models.Sponsor{
			Name:  &name,
			Image: &image,
		}

		assert.Error(sponsor.Sanitize())
	})

	t.Run("missing name, return error", func(t *testing.T) {
		assert := assert.New(t)
		uuid := uuid.New()
		image := "foo"
		sponsor := models.Sponsor{
			UUID:  &uuid,
			Image: &image,
		}

		assert.Error(sponsor.Sanitize())
	})

	t.Run("missing image, return error", func(t *testing.T) {
		assert := assert.New(t)
		uuid := uuid.New()
		name := "foo"
		sponsor := models.Sponsor{
			UUID: &uuid,
			Name: &name,
		}

		assert.Error(sponsor.Sanitize())
	})

	t.Run("give correct value, return nil", func(t *testing.T) {
		assert := assert.New(t)
		uuid := uuid.New()
		name := "foo"
		image := "bar"
		sponsor := models.Sponsor{
			UUID:  &uuid,
			Name:  &name,
			Image: &image,
		}

		assert.Nil(sponsor.Sanitize())
	})
}

func TestSponsorToAPIMeetup(t *testing.T) {
	assert := assert.New(t)
	uuid := uuid.New()
	name := "foo"
	image := "bar"
	sponsor := models.Sponsor{
		UUID:  &uuid,
		Name:  &name,
		Image: &image,
	}

	sponsorAPIModel := &api_model.Sponsor{
		UUID:  uuid.String(),
		Name:  &name,
		Image: &image,
	}
	assert.Equal(sponsorAPIModel, sponsor.ToAPIMeetup())
}
