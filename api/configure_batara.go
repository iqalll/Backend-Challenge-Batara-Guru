// This file is safe to edit. Once it exists it will not be overwritten

package api

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"Backend-Challenge-Batara-Guru/api/operations"
	"Backend-Challenge-Batara-Guru/api/operations/gift"
	"Backend-Challenge-Batara-Guru/api/operations/user"
	"Backend-Challenge-Batara-Guru/handlers"
	"Backend-Challenge-Batara-Guru/models"
	"Backend-Challenge-Batara-Guru/utils"
)

//go:generate swagger generate server --target ../../Backend-Challenge-Batara-Guru --name Batara --spec swagger.yml --server-package api --principal interface{}

func configureFlags(api *operations.BataraAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.BataraAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	mountHandler := handlers.NewHandler()
	api.BearerAuth = utils.ValidateToken
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Login
	api.UserPostLoginHandler = user.PostLoginHandlerFunc(func(plp user.PostLoginParams) middleware.Responder {
		res, err := mountHandler.Login(context.Background(), plp)
		if err != nil {
			return user.NewPostLoginBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return user.NewPostLoginOK().WithPayload(&user.PostLoginOKBody{
			Message: "Success Login",
			Data:    res,
		})
	})

	// Create User
	api.UserPostUserHandler = user.PostUserHandlerFunc(func(pup user.PostUserParams, i interface{}) middleware.Responder {
		err := mountHandler.CreateUser(context.Background(), pup)
		if err != nil {
			return user.NewPostUserBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return user.NewPostUserOK().WithPayload(&user.PostUserOKBody{
			Message: "Success Create User",
		})
	})

	// Get All Data User
	api.UserGetUserHandler = user.GetUserHandlerFunc(func(gup user.GetUserParams, i interface{}) middleware.Responder {
		res, meta, err := mountHandler.GetAllUser(context.Background(), gup)
		if err != nil {
			return user.NewGetUserBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return user.NewGetUserOK().WithPayload(&user.GetUserOKBody{
			Data:     res,
			Metadata: meta,
			Message:  "Success Get All Data User",
		})
	})

	// Update User
	api.UserPutUserIDHandler = user.PutUserIDHandlerFunc(func(pui user.PutUserIDParams, i interface{}) middleware.Responder {
		err := mountHandler.UpdateUserById(context.Background(), pui)
		if err != nil {
			return user.NewPutUserIDBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return user.NewPutUserIDOK().WithPayload(&user.PutUserIDOKBody{
			Message: "Success Update User",
		})
	})

	// Delete User
	api.UserDeleteUserIDHandler = user.DeleteUserIDHandlerFunc(func(dui user.DeleteUserIDParams, i interface{}) middleware.Responder {
		err := mountHandler.DeleteUserById(context.Background(), dui)
		if err != nil {
			return user.NewDeleteUserIDBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return user.NewDeleteUserIDOK().WithPayload(&user.DeleteUserIDOKBody{
			Message: "Success Delete User",
		})
	})

	// Create Gift
	api.GiftPostGiftHandler = gift.PostGiftHandlerFunc(func(pgp gift.PostGiftParams, i interface{}) middleware.Responder {
		err := mountHandler.CreateGift(context.Background(), pgp)
		if err != nil {
			return gift.NewPostGiftBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return gift.NewPostGiftOK().WithPayload(&gift.PostGiftOKBody{
			Message: "Success Create Gift",
		})
	})

	// Update Gift By Id
	api.GiftPutGiftIDHandler = gift.PutGiftIDHandlerFunc(func(pgi gift.PutGiftIDParams, i interface{}) middleware.Responder {
		err := mountHandler.UpdateGiftById(context.Background(), pgi)
		if err != nil {
			return gift.NewPutGiftIDBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return gift.NewPutGiftIDOK().WithPayload(&gift.PutGiftIDOKBody{
			Message: "Success Update Gift",
		})
	})

	// Delete Gift By Id
	api.GiftDeleteGiftIDHandler = gift.DeleteGiftIDHandlerFunc(func(dgi gift.DeleteGiftIDParams, i interface{}) middleware.Responder {
		err := mountHandler.DeleteGiftById(context.Background(), dgi)
		if err != nil {
			return gift.NewDeleteGiftIDBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return gift.NewDeleteGiftIDOK().WithPayload(&gift.DeleteGiftIDOKBody{
			Message: "Success Delete Gift",
		})
	})

	// Get All Gift
	api.GiftGetGiftHandler = gift.GetGiftHandlerFunc(func(ggp gift.GetGiftParams, i interface{}) middleware.Responder {
		res, meta, err := mountHandler.GetAllGift(context.Background(), ggp)
		if err != nil {
			return gift.NewGetGiftBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return gift.NewGetGiftOK().WithPayload(&gift.GetGiftOKBody{
			Data:     res,
			Metadata: meta,
			Message:  "Success Get Data All Gift",
		})
	})

	// Get Gift By Id
	api.GiftGetGiftIDHandler = gift.GetGiftIDHandlerFunc(func(ggi gift.GetGiftIDParams, i interface{}) middleware.Responder {
		res, err := mountHandler.GetGiftById(context.Background(), ggi)
		if err != nil {
			return gift.NewGetGiftIDBadRequest().WithPayload(&models.Error{Code: "400", Message: err.Error()})
		}
		return gift.NewGetGiftIDOK().WithPayload(&gift.GetGiftIDOKBody{
			Data:    res,
			Message: "Success Get Data By Id Gift",
		})
	})

	// Applies when the "Authorization" header is set
	if api.BearerAuth == nil {
		api.BearerAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// gift.PostGiftMaxParseMemory = 32 << 20
	// You may change here the memory limit for this multipart form parser. Below is the default (32 MB).
	// gift.PutGiftIDMaxParseMemory = 32 << 20

	if api.GiftDeleteGiftIDHandler == nil {
		api.GiftDeleteGiftIDHandler = gift.DeleteGiftIDHandlerFunc(func(params gift.DeleteGiftIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation gift.DeleteGiftID has not yet been implemented")
		})
	}
	if api.UserDeleteUserIDHandler == nil {
		api.UserDeleteUserIDHandler = user.DeleteUserIDHandlerFunc(func(params user.DeleteUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUserID has not yet been implemented")
		})
	}
	if api.GiftGetGiftHandler == nil {
		api.GiftGetGiftHandler = gift.GetGiftHandlerFunc(func(params gift.GetGiftParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation gift.GetGift has not yet been implemented")
		})
	}
	if api.GiftGetGiftIDHandler == nil {
		api.GiftGetGiftIDHandler = gift.GetGiftIDHandlerFunc(func(params gift.GetGiftIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation gift.GetGiftID has not yet been implemented")
		})
	}
	if api.UserGetUserHandler == nil {
		api.UserGetUserHandler = user.GetUserHandlerFunc(func(params user.GetUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		})
	}
	if api.GiftPostGiftHandler == nil {
		api.GiftPostGiftHandler = gift.PostGiftHandlerFunc(func(params gift.PostGiftParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation gift.PostGift has not yet been implemented")
		})
	}
	if api.UserPostLoginHandler == nil {
		api.UserPostLoginHandler = user.PostLoginHandlerFunc(func(params user.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PostLogin has not yet been implemented")
		})
	}
	if api.UserPostUserHandler == nil {
		api.UserPostUserHandler = user.PostUserHandlerFunc(func(params user.PostUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.PostUser has not yet been implemented")
		})
	}
	if api.GiftPutGiftIDHandler == nil {
		api.GiftPutGiftIDHandler = gift.PutGiftIDHandlerFunc(func(params gift.PutGiftIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation gift.PutGiftID has not yet been implemented")
		})
	}
	if api.UserPutUserIDHandler == nil {
		api.UserPutUserIDHandler = user.PutUserIDHandlerFunc(func(params user.PutUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.PutUserID has not yet been implemented")
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
