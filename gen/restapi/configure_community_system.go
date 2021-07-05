// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/cloudnative-id/community-system/gen/restapi/operations"
	"github.com/cloudnative-id/community-system/handlers"
	"github.com/cloudnative-id/community-system/pkg/settings"
	"github.com/cloudnative-id/community-system/pkg/storage/postgres"
)

//go:generate swagger generate server --target ../../gen --name CommunitySystem --spec ../../swagger.yaml --principal interface{}

func configureFlags(api *operations.CommunitySystemAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CommunitySystemAPI) http.Handler {
	contextLogger := log.WithFields(log.Fields{
		"handler": "PatchSponsorHandler",
	})

	api.ServeError = errors.ServeError
	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	settings := settings.NewSettings()
	store := postgres.NewPostgres(settings)

	err := store.MigrateMeetup()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	err = store.MigrateSpeaker()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	err = store.MigrateSponsor()
	if err != nil {
		contextLogger.Fatalln(err)
	}

	api.MeetupGetMeetupHandler = handlers.NewGetMeetupHandler(store)
	api.MeetupPutMeetupHandler = handlers.NewPutMeetupHandler(store)
	api.MeetupGetMeetupsHandler = handlers.NewGetMeetupsHandler(store)

	api.SponsorPutSponsorHandler = handlers.NewPutSponsorHandler(store)
	api.SponsorGetSponsorsHandler = handlers.NewGetSponsorsHandler(store)
	api.SponsorPatchSponsorHandler = handlers.NewPatchSponsorHandler(store)

	api.SpeakerGetSpeakersHandler = handlers.NewGetSpeakersHandler(store)
	api.SpeakerPutSpeakerHandler = handlers.NewPutSpeakerHandler(store)
	api.SpeakerPatchSpeakerHandler = handlers.NewPatchSpeakerHandler(store)

	api.PreServerShutdown = func() {}
	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
