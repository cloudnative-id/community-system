package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	api_speaker "github.com/cloudnative-id/community-system/gen/restapi/operations/speaker"
)

type PutSpeakerHandler struct {
	storage storage.Storage
}

func (h *PutSpeakerHandler) Handle(params api_speaker.PutSpeakerParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "PutSpeakerHandler",
	})

	meetup, exist, err := h.storage.GetMeetup(params.ID)
	if !exist {
		contextLogger.Errorf("meetup %s is not exist", params.ID)
		return api_speaker.NewPutSpeakerNotFound()
	}

	if err != nil {
		contextLogger.Errorf("getting meetup: %s", err)
		return api_speaker.NewPutSpeakerDefault(500)
	}

	speakerUUID := uuid.New()
	meetupUUID := uuid.MustParse(params.ID)

	speaker, err := models.NewSpeakerBuilder().
		SetUUID(&speakerUUID).
		SetMeetupUUID(&meetupUUID).
		SetName(params.Speaker.Name).
		SetCompany(params.Speaker.Company).
		SetPosition(params.Speaker.Position).
		SetTitle(params.Speaker.Title).
		SetImage(params.Speaker.Image).
		Build()
	if err != nil {
		contextLogger.Errorf("cannot build speaker: %s", err)
		return api_speaker.NewPutSpeakerDefault(500)
	}

	err = h.storage.WriteSpeaker(speaker)
	if err != nil {
		contextLogger.Errorf("cannot write speaker: %s", err)
		return api_speaker.NewPutSpeakerDefault(500)
	}

	meetup.Speaker = append(meetup.Speaker, speakerUUID)
	err = h.storage.WriteMeetup(meetup)
	if err != nil {
		contextLogger.Errorf("cannot write meetup: %s", err)
		return api_speaker.NewPutSpeakerDefault(500)
	}

	return api_speaker.NewPutSpeakerOK().WithPayload(&api_model.CreateObject{
		UUID:   speakerUUID.String(),
		Status: "Created",
	})
}

func NewPutSpeakerHandler(storage storage.Storage) api_speaker.PutSpeakerHandler {
	return &PutSpeakerHandler{
		storage: storage,
	}
}
