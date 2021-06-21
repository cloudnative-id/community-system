package handlers

import (
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
	var tags models.Tags
	tags = params.Meetup.Tags

	uuid := uuid.New()

	meetup, err := models.NewMeetupBuilder().
		SetUUID(&uuid).
		SetCountry(params.Meetup.Country).
		SetCity(params.Meetup.City).
		SetLocation(params.Meetup.Location).
		SetDuration(params.Meetup.Duration).
		SetTags(tags).
		SetTime(params.Meetup.Time).
		Build()
	if err != nil {
		return api_meetup.NewPutMeetupDefault(500)
	}

	err = h.storage.WriteMeetup(meetup)
	if err != nil {
		return api_meetup.NewPutMeetupDefault(500)
	}

	return api_meetup.NewPutMeetupOK().WithPayload(&api_model.CreateObject{
		UUID:   uuid.String(),
		Status: "Created",
	})
}

func NewPutmeetupHandler(storage storage.Storage) api_meetup.PutMeetupHandler {
	return &PutMeetupHandler{
		storage: storage,
	}
}
