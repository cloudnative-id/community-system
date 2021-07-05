package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"

	api_meetup "github.com/cloudnative-id/community-system/gen/restapi/operations/meetup"
)

type PatchMeetupHandler struct {
	storage storage.Storage
}

func (h *PatchMeetupHandler) Handle(params api_meetup.PatchMeetupParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "PatchMeetupHandler",
	})

	meetup, exist, err := h.storage.GetMeetup(params.MeetupID)
	if !exist {
		contextLogger.Errorf("meetup %s is not exist", params.MeetupID)
		return api_meetup.NewPatchMeetupDefault(500)
	}
	if err != nil {
		contextLogger.Errorf("cannot build meetup: %s", err)
		return api_meetup.NewPatchMeetupDefault(500)
	}

	meetup.Country = params.Meetup.Country
	meetup.City = params.Meetup.City
	meetup.Location = params.Meetup.Location
	meetup.Time = params.Meetup.Time
	meetup.Duration = params.Meetup.Duration
	meetup.Tags = params.Meetup.Tags
	meetup.RegistrationURL = params.Meetup.RegistrationURL
	meetup.Image = params.Meetup.Image

	sponsors, err := models.NewUUIDs(params.Meetup.Sponsors)
	if err != nil {
		contextLogger.Errorf("cannot parse sponsors: %s", err)
		return api_meetup.NewPatchMeetupDefault(500)
	}
	meetup.Sponsors = sponsors

	err = h.storage.WriteMeetup(meetup)
	if err != nil {
		contextLogger.Errorf("cannot write meetup: %s", err)
		return api_meetup.NewPatchMeetupDefault(500)
	}

	return api_meetup.NewPatchMeetupNoContent()
}

func NewPatchMeetupHandler(storage storage.Storage) api_meetup.PatchMeetupHandler {
	return &PatchMeetupHandler{
		storage: storage,
	}
}
