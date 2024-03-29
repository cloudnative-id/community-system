package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	api_meetup "github.com/cloudnative-id/community-system/gen/restapi/operations/meetup"
)

type GetMeetupsHandler struct {
	storage storage.Storage
}

func (h *GetMeetupsHandler) Handle(params api_meetup.GetMeetupsParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "GetMeetupsHandler",
	})

	meetups, exist, err := h.storage.GetMeetups()

	if !exist {
		contextLogger.Errorf("meetups is not exist")
		return api_meetup.NewGetMeetupsNotFound()
	}

	if err != nil {
		contextLogger.Errorf("getting meetups: %s", err)
		return api_meetup.NewGetMeetupsDefault(500)
	}

	var meetupsResponse []*api_model.Meetup
	for _, meetup := range meetups {
		meetupsResponse = append(meetupsResponse, meetup.ToAPIMeetup())
	}

	return api_meetup.NewGetMeetupsOK().WithPayload(meetupsResponse)
}

func NewGetMeetupsHandler(storage storage.Storage) api_meetup.GetMeetupsHandler {
	return &GetMeetupsHandler{
		storage: storage,
	}
}
