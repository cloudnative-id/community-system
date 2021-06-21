// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cloudnative-id/community-system/gen/restapi/operations"
	"github.com/cloudnative-id/community-system/gen/restapi/operations/meetup"
	"github.com/cloudnative-id/community-system/handlers"
	"github.com/cloudnative-id/community-system/pkg/settings"
	"github.com/cloudnative-id/community-system/pkg/storage/postgres"
)

//go:generate swagger generate server --target ../../gen --name CommunitySystem --spec ../../swagger.yaml --principal interface{}

func configureFlags(api *operations.CommunitySystemAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CommunitySystemAPI) http.Handler {
	api.ServeError = errors.ServeError
	api.UseSwaggerUI()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	settings := settings.NewSettings()
	store := postgres.NewPostgres(settings)

	err := store.MigrateMeetup()
	if err != nil {
		log.Fatalln(err)
	}

	err = store.MigrateSpeaker()
	if err != nil {
		log.Fatalln(err)
	}

	err = store.MigrateSponsor()
	if err != nil {
		log.Fatalln(err)
	}

	api.MeetupGetMeetupHandler = handlers.NewGetmeetupHandler(store)
	api.MeetupPutMeetupHandler = handlers.NewPutmeetupHandler(store)
	api.MeetupGetMeetupsHandler = handlers.NewGetmeetupsHandler(store)

	if api.MeetupGetMeetupHandler == nil {
		api.MeetupGetMeetupHandler = meetup.GetMeetupHandlerFunc(func(params meetup.GetMeetupParams) middleware.Responder {
			return middleware.NotImplemented("operation meetup.GetMeetup has not yet been implemented")
		})
	}
	if api.MeetupGetMeetupsHandler == nil {
		api.MeetupGetMeetupsHandler = meetup.GetMeetupsHandlerFunc(func(params meetup.GetMeetupsParams) middleware.Responder {
			return middleware.NotImplemented("operation meetup.GetMeetups has not yet been implemented")
		})
	}
	if api.MeetupPutMeetupHandler == nil {
		api.MeetupPutMeetupHandler = meetup.PutMeetupHandlerFunc(func(params meetup.PutMeetupParams) middleware.Responder {
			return middleware.NotImplemented("operation meetup.PutMeetup has not yet been implemented")
		})
	}

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
