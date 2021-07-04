package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	api_speaker "github.com/cloudnative-id/community-system/gen/restapi/operations/speaker"
)

type GetSpeakersHandler struct {
	storage storage.Storage
}

func (h *GetSpeakersHandler) Handle(params api_speaker.GetSpeakersParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "GetSpeakersHandler",
	})

	speakers, exist, err := h.storage.GetSpeakers(params.MeetupID)

	if !exist {
		contextLogger.Errorf("speaker is not exist")
		return api_speaker.NewGetSpeakersNotFound()
	}

	if err != nil {
		contextLogger.Errorf("getting speaker: %s", err)
		return api_speaker.NewGetSpeakersDefault(500)
	}

	var speakersResponse []*api_model.Speaker
	for _, speaker := range speakers {
		speakersResponse = append(speakersResponse, speaker.ToAPIMeetup())
	}

	return api_speaker.NewGetSpeakersOK().WithPayload(speakersResponse)
}

func NewGetSpeakersHandler(storage storage.Storage) api_speaker.GetSpeakersHandler {
	return &GetSpeakersHandler{
		storage: storage,
	}
}
