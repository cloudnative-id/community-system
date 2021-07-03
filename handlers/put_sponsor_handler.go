package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/models"
	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"
	"github.com/google/uuid"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	api_sponsor "github.com/cloudnative-id/community-system/gen/restapi/operations/sponsor"
)

type PutSponsorHandler struct {
	storage storage.Storage
}

func (h *PutSponsorHandler) Handle(params api_sponsor.PutSponsorParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "PutSponsorHandler",
	})

	sponsorUUID := uuid.New()

	sponsor, err := models.NewSponsorBuilder().
		SetUUID(&sponsorUUID).
		SetName(params.Sponsor.Name).
		SetImage(params.Sponsor.Image).
		Build()
	if err != nil {
		contextLogger.Errorf("cannot build sponsor: %s", err)
		return api_sponsor.NewPutSponsorDefault(500)
	}

	err = h.storage.WriteSponsor(sponsor)
	if err != nil {
		contextLogger.Errorf("cannot write sponsor: %s", err)
		return api_sponsor.NewPutSponsorDefault(500)
	}

	return api_sponsor.NewPutSponsorOK().WithPayload(&api_model.CreateObject{
		UUID:   sponsorUUID.String(),
		Status: "Created",
	})
}

func NewPutSponsorHandler(storage storage.Storage) api_sponsor.PutSponsorHandler {
	return &PutSponsorHandler{
		storage: storage,
	}
}
