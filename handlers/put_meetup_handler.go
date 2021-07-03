package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	api_meetup "github.com/cloudnative-id/community-system/gen/restapi/operations/meetup"
)

type PutMeetupHandler struct {
	storage storage.Storage
}

func (h *PutMeetupHandler) Handle(params api_meetup.PutMeetupParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "PutMeetupHandler",
	})

	var tags models.Tags
	tags = params.Meetup.Tags

	MeetupUUID := uuid.New()

	meetup, err := models.NewMeetupBuilder().
		SetUUID(&MeetupUUID).
		SetCountry(params.Meetup.Country).
		SetCity(params.Meetup.City).
		SetLocation(params.Meetup.Location).
		SetDuration(params.Meetup.Duration).
		SetTags(tags).
		SetTime(params.Meetup.Time).
		Build()
	if err != nil {
		contextLogger.Errorf("cannot build meetup: %s", err)
		return api_meetup.NewPutMeetupDefault(500)
	}

	err = h.storage.WriteMeetup(meetup)
	if err != nil {
		contextLogger.Errorf("cannot write meetup: %s", err)
		return api_meetup.NewPutMeetupDefault(500)
	}

	return api_meetup.NewPutMeetupOK().WithPayload(&api_model.CreateObject{
		UUID:   MeetupUUID.String(),
		Status: "Created",
	})
}

func NewPutMeetupHandler(storage storage.Storage) api_meetup.PutMeetupHandler {
	return &PutMeetupHandler{
		storage: storage,
	}
}
