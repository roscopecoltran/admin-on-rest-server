// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/apply_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/authentication_rest_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/data_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/data_source_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/permission_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/role_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/schema_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/user_controller"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target ../server --name EasyAdmin --spec ../swagger.yaml

func configureFlags(api *operations.EasyAdminAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.EasyAdminAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.DataSourceControllerAddDataSourceUsingPOSTHandler = data_source_controller.AddDataSourceUsingPOSTHandlerFunc(func(params data_source_controller.AddDataSourceUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation data_source_controller.AddDataSourceUsingPOST has not yet been implemented")
	})
	api.SchemaControllerAddEntityUsingPOSTHandler = schema_controller.AddEntityUsingPOSTHandlerFunc(func(params schema_controller.AddEntityUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.AddEntityUsingPOST has not yet been implemented")
	})
	api.SchemaControllerAddFieldUsingPOSTHandler = schema_controller.AddFieldUsingPOSTHandlerFunc(func(params schema_controller.AddFieldUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.AddFieldUsingPOST has not yet been implemented")
	})
	api.PermissionControllerAddPermissionUsingPOSTHandler = permission_controller.AddPermissionUsingPOSTHandlerFunc(func(params permission_controller.AddPermissionUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation permission_controller.AddPermissionUsingPOST has not yet been implemented")
	})
	api.RoleControllerAddRoleUsingPOSTHandler = role_controller.AddRoleUsingPOSTHandlerFunc(func(params role_controller.AddRoleUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation role_controller.AddRoleUsingPOST has not yet been implemented")
	})
	api.UserControllerAddUserUsingPOSTHandler = user_controller.AddUserUsingPOSTHandlerFunc(func(params user_controller.AddUserUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation user_controller.AddUserUsingPOST has not yet been implemented")
	})
	api.ApplyControllerApplyUsingPOSTHandler = apply_controller.ApplyUsingPOSTHandlerFunc(func(params apply_controller.ApplyUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation apply_controller.ApplyUsingPOST has not yet been implemented")
	})
	api.AuthenticationRestControllerCreateAuthenticationTokenUsingPOSTHandler = authentication_rest_controller.CreateAuthenticationTokenUsingPOSTHandlerFunc(func(params authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation authentication_rest_controller.CreateAuthenticationTokenUsingPOST has not yet been implemented")
	})
	api.DataControllerDataMutationUsingDELETEHandler = data_controller.DataMutationUsingDELETEHandlerFunc(func(params data_controller.DataMutationUsingDELETEParams) middleware.Responder {
		return middleware.NotImplemented("operation data_controller.DataMutationUsingDELETE has not yet been implemented")
	})
	api.DataControllerDataMutationUsingPOSTHandler = data_controller.DataMutationUsingPOSTHandlerFunc(func(params data_controller.DataMutationUsingPOSTParams) middleware.Responder {
		return middleware.NotImplemented("operation data_controller.DataMutationUsingPOST has not yet been implemented")
	})
	api.DataControllerDataMutationUsingPUTHandler = data_controller.DataMutationUsingPUTHandlerFunc(func(params data_controller.DataMutationUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation data_controller.DataMutationUsingPUT has not yet been implemented")
	})
	api.DataControllerDataQueryUsingGETHandler = data_controller.DataQueryUsingGETHandlerFunc(func(params data_controller.DataQueryUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation data_controller.DataQueryUsingGET has not yet been implemented")
	})
	api.DataSourceControllerEditDataSourceUsingPUTHandler = data_source_controller.EditDataSourceUsingPUTHandlerFunc(func(params data_source_controller.EditDataSourceUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation data_source_controller.EditDataSourceUsingPUT has not yet been implemented")
	})
	api.SchemaControllerEditEntityUsingPUTHandler = schema_controller.EditEntityUsingPUTHandlerFunc(func(params schema_controller.EditEntityUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.EditEntityUsingPUT has not yet been implemented")
	})
	api.SchemaControllerEditFieldUsingPUTHandler = schema_controller.EditFieldUsingPUTHandlerFunc(func(params schema_controller.EditFieldUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.EditFieldUsingPUT has not yet been implemented")
	})
	api.UserControllerEditFieldUsingPUT1Handler = user_controller.EditFieldUsingPUT1HandlerFunc(func(params user_controller.EditFieldUsingPUT1Params) middleware.Responder {
		return middleware.NotImplemented("operation user_controller.EditFieldUsingPUT1 has not yet been implemented")
	})
	api.PermissionControllerEditPermissionUsingPUTHandler = permission_controller.EditPermissionUsingPUTHandlerFunc(func(params permission_controller.EditPermissionUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation permission_controller.EditPermissionUsingPUT has not yet been implemented")
	})
	api.RoleControllerEditRoleUsingPUTHandler = role_controller.EditRoleUsingPUTHandlerFunc(func(params role_controller.EditRoleUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation role_controller.EditRoleUsingPUT has not yet been implemented")
	})
	api.SchemaControllerFindAllFieldsUsingGETHandler = schema_controller.FindAllFieldsUsingGETHandlerFunc(func(params schema_controller.FindAllFieldsUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.FindAllFieldsUsingGET has not yet been implemented")
	})
	api.SchemaControllerFindEntityFieldsUsingGETHandler = schema_controller.FindEntityFieldsUsingGETHandlerFunc(func(params schema_controller.FindEntityFieldsUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.FindEntityFieldsUsingGET has not yet been implemented")
	})
	api.SchemaControllerFindOneFieldUsingGETHandler = schema_controller.FindOneFieldUsingGETHandlerFunc(func(params schema_controller.FindOneFieldUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.FindOneFieldUsingGET has not yet been implemented")
	})
	api.DataControllerFindOneUsingGETHandler = data_controller.FindOneUsingGETHandlerFunc(func(params data_controller.FindOneUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation data_controller.FindOneUsingGET has not yet been implemented")
	})
	api.SchemaControllerFindOneUsingGET1Handler = schema_controller.FindOneUsingGET1HandlerFunc(func(params schema_controller.FindOneUsingGET1Params) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.FindOneUsingGET1 has not yet been implemented")
	})
	api.DataSourceControllerFindRoleUsingGETHandler = data_source_controller.FindRoleUsingGETHandlerFunc(func(params data_source_controller.FindRoleUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation data_source_controller.FindRoleUsingGET has not yet been implemented")
	})
	api.RoleControllerFindRoleUsingGET1Handler = role_controller.FindRoleUsingGET1HandlerFunc(func(params role_controller.FindRoleUsingGET1Params) middleware.Responder {
		return middleware.NotImplemented("operation role_controller.FindRoleUsingGET1 has not yet been implemented")
	})
	api.PermissionControllerFindUserUsingGETHandler = permission_controller.FindUserUsingGETHandlerFunc(func(params permission_controller.FindUserUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation permission_controller.FindUserUsingGET has not yet been implemented")
	})
	api.UserControllerFindUserUsingGET1Handler = user_controller.FindUserUsingGET1HandlerFunc(func(params user_controller.FindUserUsingGET1Params) middleware.Responder {
		return middleware.NotImplemented("operation user_controller.FindUserUsingGET1 has not yet been implemented")
	})
	api.UserControllerGetAuthenticatedUserUsingGETHandler = user_controller.GetAuthenticatedUserUsingGETHandlerFunc(func(params user_controller.GetAuthenticatedUserUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation user_controller.GetAuthenticatedUserUsingGET has not yet been implemented")
	})
	api.SchemaControllerGetSchemasUsingGETHandler = schema_controller.GetSchemasUsingGETHandlerFunc(func(params schema_controller.GetSchemasUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.GetSchemasUsingGET has not yet been implemented")
	})
	api.DataSourceControllerListUsingGETHandler = data_source_controller.ListUsingGETHandlerFunc(func(params data_source_controller.ListUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation data_source_controller.ListUsingGET has not yet been implemented")
	})
	api.PermissionControllerListUsingGET1Handler = permission_controller.ListUsingGET1HandlerFunc(func(params permission_controller.ListUsingGET1Params) middleware.Responder {
		return middleware.NotImplemented("operation permission_controller.ListUsingGET1 has not yet been implemented")
	})
	api.RoleControllerListUsingGET2Handler = role_controller.ListUsingGET2HandlerFunc(func(params role_controller.ListUsingGET2Params) middleware.Responder {
		return middleware.NotImplemented("operation role_controller.ListUsingGET2 has not yet been implemented")
	})
	api.UserControllerListUsingGET3Handler = user_controller.ListUsingGET3HandlerFunc(func(params user_controller.ListUsingGET3Params) middleware.Responder {
		return middleware.NotImplemented("operation user_controller.ListUsingGET3 has not yet been implemented")
	})
	api.AuthenticationRestControllerRefreshAndGetAuthenticationTokenUsingGETHandler = authentication_rest_controller.RefreshAndGetAuthenticationTokenUsingGETHandlerFunc(func(params authentication_rest_controller.RefreshAndGetAuthenticationTokenUsingGETParams) middleware.Responder {
		return middleware.NotImplemented("operation authentication_rest_controller.RefreshAndGetAuthenticationTokenUsingGET has not yet been implemented")
	})
	api.SchemaControllerResetCurrentDsUsingPUTHandler = schema_controller.ResetCurrentDsUsingPUTHandlerFunc(func(params schema_controller.ResetCurrentDsUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.ResetCurrentDsUsingPUT has not yet been implemented")
	})
	api.SchemaControllerSyncSchemasUsingPUTHandler = schema_controller.SyncSchemasUsingPUTHandlerFunc(func(params schema_controller.SyncSchemasUsingPUTParams) middleware.Responder {
		return middleware.NotImplemented("operation schema_controller.SyncSchemasUsingPUT has not yet been implemented")
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
