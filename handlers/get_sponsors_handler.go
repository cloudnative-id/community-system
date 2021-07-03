package handlers

import (
	log "github.com/sirupsen/logrus"

	"github.com/cloudnative-id/community-system/pkg/storage"
	"github.com/go-openapi/runtime/middleware"

	api_model "github.com/cloudnative-id/community-system/gen/models"
	api_sponsor "github.com/cloudnative-id/community-system/gen/restapi/operations/sponsor"
)

type GetSponsorsHandler struct {
	storage storage.Storage
}

func (h *GetSponsorsHandler) Handle(params api_sponsor.GetSponsorsParams) middleware.Responder {
	contextLogger := log.WithFields(log.Fields{
		"handler": "GetSponsorsHandler",
	})

	sponsors, exist, err := h.storage.GetSponsors()

	if !exist {
		contextLogger.Errorf("sponsor is not exist")
		return api_sponsor.NewGetSponsorsNotFound()
	}

	if err != nil {
		contextLogger.Errorf("getting sponsor: %s", err)
		return api_sponsor.NewGetSponsorsDefault(500)
	}

	var sponsorsResponse []*api_model.Sponsor
	for _, sponsor := range sponsors {
		sponsorsResponse = append(sponsorsResponse, sponsor.ToAPIMeetup())
	}

	return api_sponsor.NewGetSponsorsOK().WithPayload(sponsorsResponse)
}

func NewGetSponsorsHandler(storage storage.Storage) api_sponsor.GetSponsorsHandler {
	return &GetSponsorsHandler{
		storage: storage,
	}
}
