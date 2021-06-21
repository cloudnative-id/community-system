package handlers

import (
	api_model "github.com/cloudnative-id/community-system/gen/models"
	api_meetup "github.com/cloudnative-id/community-system/gen/restapi/operations/meetup"
	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"
)

type GetMeetupHandler struct {
	storage storage.Storage
}

func (h *GetMeetupHandler) Handle(params api_meetup.GetMeetupParams) middleware.Responder {
	meetup, exist, err := h.storage.GetMeetup(params.ID)

	if !exist {
		return api_meetup.NewGetMeetupNotFound()
	}

	if err != nil {
		return api_meetup.NewGetMeetupDefault(500)
	}

	return api_meetup.NewGetMeetupOK().WithPayload(&api_model.Meetup{
		UUID:            meetup.UUID.String(),
		Country:         meetup.Country,
		City:            meetup.City,
		Location:        meetup.Location,
		Tags:            meetup.Tags,
		Time:            meetup.Time,
		Duration:        meetup.Duration,
		RegistrationURL: meetup.RegistrationURL,
		Image:           meetup.Image,
		Speaker:         meetup.Speaker.ArrayString(),
		Sponsors:        meetup.Sponsors.ArrayString(),
		Status:          meetup.Status,
	})
}

func NewGetmeetupHandler(storage storage.Storage) api_meetup.GetMeetupHandler {
	return &GetMeetupHandler{
		storage: storage,
	}
}
