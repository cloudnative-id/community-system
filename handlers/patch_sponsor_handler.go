package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"

	api_sponsor "github.com/cloudnative-id/community-system/gen/restapi/operations/sponsor"
)

type PatchSponsorHandler struct {
	storage storage.Storage
}

func (h *PatchSponsorHandler) Handle(params api_sponsor.PatchSponsorParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "PatchSponsorHandler",
	})

	sponsor, exist, err := h.storage.GetSponsor(params.SponsorID)
	if !exist {
		contextLogger.Errorf("sponsor %s is not exist", params.SponsorID)
		return api_sponsor.NewPatchSponsorDefault(500)
	}
	if err != nil {
		contextLogger.Errorf("cannot build sponsor: %s", err)
		return api_sponsor.NewPatchSponsorDefault(500)
	}

	sponsor.Image = params.Sponsor.Image

	err = h.storage.WriteSponsor(sponsor)
	if err != nil {
		contextLogger.Errorf("cannot write sponsor: %s", err)
		return api_sponsor.NewPatchSponsorDefault(500)
	}

	return api_sponsor.NewPatchSponsorNoContent()
}

func NewPatchSponsorHandler(storage storage.Storage) api_sponsor.PatchSponsorHandler {
	return &PatchSponsorHandler{
		storage: storage,
	}
}
