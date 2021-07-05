package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"

	api_speaker "github.com/cloudnative-id/community-system/gen/restapi/operations/speaker"
)

type PatchSpeakerHandler struct {
	storage storage.Storage
}

func (h *PatchSpeakerHandler) Handle(params api_speaker.PatchSpeakerParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "PatchSpeakerHandler",
	})

	speaker, exist, err := h.storage.GetSpeaker(params.MeetupID, params.SpeakerID)
	if !exist {
		contextLogger.Errorf("speaker %s is not exist", params.SpeakerID)
		return api_speaker.NewPatchSpeakerDefault(500)
	}
	if err != nil {
		contextLogger.Errorf("cannot build speaker: %s", err)
		return api_speaker.NewPatchSpeakerDefault(500)
	}

	speaker.Name = params.Speaker.Name
	speaker.Position = params.Speaker.Position
	speaker.Company = params.Speaker.Company
	speaker.Title = params.Speaker.Title
	speaker.Image = params.Speaker.Image

	err = h.storage.WriteSpeaker(speaker)
	if err != nil {
		contextLogger.Errorf("cannot write speaker: %s", err)
		return api_speaker.NewPatchSpeakerDefault(500)
	}

	return api_speaker.NewPatchSpeakerNoContent()
}

func NewPatchSpeakerHandler(storage storage.Storage) api_speaker.PatchSpeakerHandler {
	return &PatchSpeakerHandler{
		storage: storage,
	}
}
