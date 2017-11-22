package restapi

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/mikkeloscar/gin-swagger/api"
	"github.com/mikkeloscar/gin-swagger/middleware"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/apply_controller"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/authentication_rest_controller"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/data_controller"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/data_source_controller"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/permission_controller"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/role_controller"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/schema_controller"
	"github.com/roscopecoltran/admin-on-rest-server/.experimental/gin-swagger/restapi/operations/user_controller"
	log "github.com/sirupsen/logrus"
	ginoauth2 "github.com/zalando/gin-oauth2"
	"golang.org/x/oauth2"
)

// Routes defines all the routes of the API service.
type Routes struct {
	*gin.Engine
	AddDataSourceUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddEntityUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddFieldUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddPermissionUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddRoleUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	AddUserUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ApplyUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	CreateAuthenticationTokenUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingDELETE struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingPOST struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataMutationUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	DataQueryUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditDataSourceUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditEntityUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditFieldUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditFieldUsingPUT1 struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditPermissionUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	EditRoleUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindAllFieldsUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindEntityFieldsUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindOneFieldUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindOneUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindOneUsingGET1 struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindRoleUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindRoleUsingGET1 struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindUserUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	FindUserUsingGET1 struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	GetAuthenticatedUserUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	GetSchemasUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ListUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ListUsingGET1 struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ListUsingGET2 struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ListUsingGET3 struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	RefreshAndGetAuthenticationTokenUsingGET struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	ResetCurrentDsUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
	SyncSchemasUsingPUT struct {
		*gin.RouterGroup
		Auth gin.HandlerFunc
		Post *gin.RouterGroup
	}
}

// configureWellKnown enables and configures /.well-known endpoints.
func (r *Routes) configureWellKnown(healthFunc func() bool) {
	wellKnown := r.Group("/.well-known")
	{
		wellKnown.GET("/schema-discovery", func(ctx *gin.Context) {
			discovery := struct {
				SchemaURL  string `json:"schema_url"`
				SchemaType string `json:"schema_type"`
				UIURL      string `json:"ui_url"`
			}{
				SchemaURL:  "/swagger.json",
				SchemaType: "swagger-2.0",
			}
			ctx.JSON(http.StatusOK, &discovery)
		})
		wellKnown.GET("/health", healthHandler(healthFunc))
	}

	r.GET("/swagger.json", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, string(SwaggerJSON))
	})
}

// healthHandler is the health HTTP handler used for the /.well-known/health
// route if enabled.
func healthHandler(healthFunc func() bool) gin.HandlerFunc {
	healthy := healthFunc
	return func(ctx *gin.Context) {
		health := struct {
			Health bool `json:"health"`
		}{
			Health: healthy(),
		}

		if !health.Health {
			ctx.JSON(http.StatusServiceUnavailable, &health)
		} else {
			ctx.JSON(http.StatusOK, &health)
		}
	}
}

/*
func infoHandler(infoFunc func() bool) gin.HandlerFunc {
	info := infoFunc
	return func(ctx *gin.Context) {
		info := struct {
			Infos bool `json:"health"`
		}{
			Infos: info(),
		}

		if !info.Infos {
			ctx.JSON(http.StatusServiceUnavailable, &info)
		} else {
			ctx.JSON(http.StatusOK, &info)
		}
	}
}
*/

// Service is the interface that must be implemented in order to provide
// business logic for the API service.
type Service interface {
	Healthy() bool
	Info(ctx *gin.Context) *api.Response
	AddDataSourceUsingPOST(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingPOSTParams) *api.Response
	AddEntityUsingPOST(ctx *gin.Context, params *schema_controller.AddEntityUsingPOSTParams) *api.Response
	AddFieldUsingPOST(ctx *gin.Context, params *schema_controller.AddFieldUsingPOSTParams) *api.Response
	AddPermissionUsingPOST(ctx *gin.Context, params *permission_controller.AddPermissionUsingPOSTParams) *api.Response
	AddRoleUsingPOST(ctx *gin.Context, params *role_controller.AddRoleUsingPOSTParams) *api.Response
	AddUserUsingPOST(ctx *gin.Context, params *user_controller.AddUserUsingPOSTParams) *api.Response
	ApplyUsingPOST(ctx *gin.Context, params *apply_controller.ApplyUsingPOSTParams) *api.Response
	CreateAuthenticationTokenUsingPOST(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) *api.Response
	DataMutationUsingDELETE(ctx *gin.Context, params *data_controller.DataMutationUsingDELETEParams) *api.Response
	DataMutationUsingPOST(ctx *gin.Context, params *data_controller.DataMutationUsingPOSTParams) *api.Response
	DataMutationUsingPUT(ctx *gin.Context, params *data_controller.DataMutationUsingPUTParams) *api.Response
	DataQueryUsingGET(ctx *gin.Context, params *data_controller.DataQueryUsingGETParams) *api.Response
	EditDataSourceUsingPUT(ctx *gin.Context, params *data_source_controller.EditDataSourceUsingPUTParams) *api.Response
	EditEntityUsingPUT(ctx *gin.Context, params *schema_controller.EditEntityUsingPUTParams) *api.Response
	EditFieldUsingPUT(ctx *gin.Context, params *schema_controller.EditFieldUsingPUTParams) *api.Response
	EditFieldUsingPUT1(ctx *gin.Context, params *user_controller.EditFieldUsingPUT1Params) *api.Response
	EditPermissionUsingPUT(ctx *gin.Context, params *permission_controller.EditPermissionUsingPUTParams) *api.Response
	EditRoleUsingPUT(ctx *gin.Context, params *role_controller.EditRoleUsingPUTParams) *api.Response
	FindAllFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindAllFieldsUsingGETParams) *api.Response
	FindEntityFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindEntityFieldsUsingGETParams) *api.Response
	FindOneFieldUsingGET(ctx *gin.Context, params *schema_controller.FindOneFieldUsingGETParams) *api.Response
	FindOneUsingGET(ctx *gin.Context, params *data_controller.FindOneUsingGETParams) *api.Response
	FindOneUsingGET1(ctx *gin.Context, params *schema_controller.FindOneUsingGET1Params) *api.Response
	FindRoleUsingGET(ctx *gin.Context, params *data_source_controller.FindRoleUsingGETParams) *api.Response
	FindRoleUsingGET1(ctx *gin.Context, params *role_controller.FindRoleUsingGET1Params) *api.Response
	FindUserUsingGET(ctx *gin.Context, params *permission_controller.FindUserUsingGETParams) *api.Response
	FindUserUsingGET1(ctx *gin.Context, params *user_controller.FindUserUsingGET1Params) *api.Response
	GetAuthenticatedUserUsingGET(ctx *gin.Context) *api.Response
	GetSchemasUsingGET(ctx *gin.Context) *api.Response
	ListUsingGET(ctx *gin.Context) *api.Response
	ListUsingGET1(ctx *gin.Context, params *permission_controller.ListUsingGET1Params) *api.Response
	ListUsingGET2(ctx *gin.Context) *api.Response
	ListUsingGET3(ctx *gin.Context) *api.Response
	RefreshAndGetAuthenticationTokenUsingGET(ctx *gin.Context) *api.Response
	ResetCurrentDsUsingPUT(ctx *gin.Context, params *schema_controller.ResetCurrentDsUsingPUTParams) *api.Response
	SyncSchemasUsingPUT(ctx *gin.Context, params *schema_controller.SyncSchemasUsingPUTParams) *api.Response
}

func ginizePath(path string) string {
	return strings.Replace(strings.Replace(path, "{", ":", -1), "}", "", -1)
}

// configureRoutes configures the routes for the API service.
func configureRoutes(service Service, enableAuth bool, tokenURL string) *Routes {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(middleware.LogrusLogger())
	routes := &Routes{Engine: engine}

	routes.AddDataSourceUsingPOST.RouterGroup = routes.Group("")
	routes.AddDataSourceUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	if enableAuth {
		routeTokenURL := tokenURL
		if routeTokenURL == "" {
			routeTokenURL = "http://info.services.auth.localhost/oauth2/tokeninfo"
		}
		routes.AddDataSourceUsingPOST.Auth = ginoauth2.Auth(
			middleware.ScopesAuth("uid"),
			oauth2.Endpoint{
				TokenURL: routeTokenURL,
			},
		)
		routes.AddDataSourceUsingPOST.RouterGroup.Use(routes.AddDataSourceUsingPOST.Auth)
	}
	routes.AddDataSourceUsingPOST.Post = routes.AddDataSourceUsingPOST.Group("")
	routes.AddDataSourceUsingPOST.Post.POST(ginizePath("/datasource/_datasource"), data_source_controller.BusinessLogicAddDataSourceUsingPOST(service.AddDataSourceUsingPOST))

	routes.AddEntityUsingPOST.RouterGroup = routes.Group("")
	routes.AddEntityUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))

	routes.AddEntityUsingPOST.Post = routes.AddEntityUsingPOST.Group("")
	routes.AddEntityUsingPOST.Post.POST(ginizePath("/schemas/_entitys"), schema_controller.BusinessLogicAddEntityUsingPOST(service.AddEntityUsingPOST))

	routes.AddFieldUsingPOST.RouterGroup = routes.Group("")
	routes.AddFieldUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.AddFieldUsingPOST.Post = routes.AddFieldUsingPOST.Group("")
	routes.AddFieldUsingPOST.Post.POST(ginizePath("/schemas/_fields"), schema_controller.BusinessLogicAddFieldUsingPOST(service.AddFieldUsingPOST))

	routes.AddPermissionUsingPOST.RouterGroup = routes.Group("")
	routes.AddPermissionUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.AddPermissionUsingPOST.Post = routes.AddPermissionUsingPOST.Group("")
	routes.AddPermissionUsingPOST.Post.POST(ginizePath("/permission/_permission"), permission_controller.BusinessLogicAddPermissionUsingPOST(service.AddPermissionUsingPOST))

	routes.AddRoleUsingPOST.RouterGroup = routes.Group("")
	routes.AddRoleUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.AddRoleUsingPOST.Post = routes.AddRoleUsingPOST.Group("")
	routes.AddRoleUsingPOST.Post.POST(ginizePath("/role/_roles"), role_controller.BusinessLogicAddRoleUsingPOST(service.AddRoleUsingPOST))

	routes.AddUserUsingPOST.RouterGroup = routes.Group("")
	routes.AddUserUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.AddUserUsingPOST.Post = routes.AddUserUsingPOST.Group("")
	routes.AddUserUsingPOST.Post.POST(ginizePath("/user/_users"), user_controller.BusinessLogicAddUserUsingPOST(service.AddUserUsingPOST))

	routes.ApplyUsingPOST.RouterGroup = routes.Group("")
	routes.ApplyUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.ApplyUsingPOST.Post = routes.ApplyUsingPOST.Group("")
	routes.ApplyUsingPOST.Post.POST(ginizePath("/apply"), apply_controller.BusinessLogicApplyUsingPOST(service.ApplyUsingPOST))

	routes.CreateAuthenticationTokenUsingPOST.RouterGroup = routes.Group("")
	routes.CreateAuthenticationTokenUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.CreateAuthenticationTokenUsingPOST.Post = routes.CreateAuthenticationTokenUsingPOST.Group("")
	routes.CreateAuthenticationTokenUsingPOST.Post.POST(ginizePath("/auth"), authentication_rest_controller.BusinessLogicCreateAuthenticationTokenUsingPOST(service.CreateAuthenticationTokenUsingPOST))

	routes.DataMutationUsingDELETE.RouterGroup = routes.Group("")
	routes.DataMutationUsingDELETE.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.DataMutationUsingDELETE.Post = routes.DataMutationUsingDELETE.Group("")
	routes.DataMutationUsingDELETE.Post.DELETE(ginizePath("/api/{entity}/{id}"), data_controller.BusinessLogicDataMutationUsingDELETE(service.DataMutationUsingDELETE))

	routes.DataMutationUsingPOST.RouterGroup = routes.Group("")
	routes.DataMutationUsingPOST.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.DataMutationUsingPOST.Post = routes.DataMutationUsingPOST.Group("")
	routes.DataMutationUsingPOST.Post.POST(ginizePath("/api/{entity}"), data_controller.BusinessLogicDataMutationUsingPOST(service.DataMutationUsingPOST))

	routes.DataMutationUsingPUT.RouterGroup = routes.Group("")
	routes.DataMutationUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.DataMutationUsingPUT.Post = routes.DataMutationUsingPUT.Group("")
	routes.DataMutationUsingPUT.Post.PUT(ginizePath("/api/{entity}/{id}"), data_controller.BusinessLogicDataMutationUsingPUT(service.DataMutationUsingPUT))

	routes.DataQueryUsingGET.RouterGroup = routes.Group("")
	routes.DataQueryUsingGET.Post = routes.DataQueryUsingGET.Group("")
	routes.DataQueryUsingGET.Post.GET(ginizePath("/api/{entity}"), data_controller.BusinessLogicDataQueryUsingGET(service.DataQueryUsingGET))

	routes.EditDataSourceUsingPUT.RouterGroup = routes.Group("")
	routes.EditDataSourceUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.EditDataSourceUsingPUT.Post = routes.EditDataSourceUsingPUT.Group("")
	routes.EditDataSourceUsingPUT.Post.PUT(ginizePath("/datasource/_datasource/put/{id}"), data_source_controller.BusinessLogicEditDataSourceUsingPUT(service.EditDataSourceUsingPUT))

	routes.EditEntityUsingPUT.RouterGroup = routes.Group("")
	routes.EditEntityUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.EditEntityUsingPUT.Post = routes.EditEntityUsingPUT.Group("")
	routes.EditEntityUsingPUT.Post.PUT(ginizePath("/schemas/_entitys/put/{id}"), schema_controller.BusinessLogicEditEntityUsingPUT(service.EditEntityUsingPUT))

	routes.EditFieldUsingPUT.RouterGroup = routes.Group("")
	routes.EditFieldUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.EditFieldUsingPUT.Post = routes.EditFieldUsingPUT.Group("")
	routes.EditFieldUsingPUT.Post.PUT(ginizePath("/schemas/_fields/put/{id}"), schema_controller.BusinessLogicEditFieldUsingPUT(service.EditFieldUsingPUT))

	routes.EditFieldUsingPUT1.RouterGroup = routes.Group("")
	routes.EditFieldUsingPUT1.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.EditFieldUsingPUT1.Post = routes.EditFieldUsingPUT1.Group("")
	routes.EditFieldUsingPUT1.Post.PUT(ginizePath("/user/_users/put/{id}"), user_controller.BusinessLogicEditFieldUsingPUT1(service.EditFieldUsingPUT1))

	routes.EditPermissionUsingPUT.RouterGroup = routes.Group("")
	routes.EditPermissionUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.EditPermissionUsingPUT.Post = routes.EditPermissionUsingPUT.Group("")
	routes.EditPermissionUsingPUT.Post.PUT(ginizePath("/permission/_permission/{id}"), permission_controller.BusinessLogicEditPermissionUsingPUT(service.EditPermissionUsingPUT))

	routes.EditRoleUsingPUT.RouterGroup = routes.Group("")
	routes.EditRoleUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.EditRoleUsingPUT.Post = routes.EditRoleUsingPUT.Group("")
	routes.EditRoleUsingPUT.Post.PUT(ginizePath("/role/_roles/put/{id}"), role_controller.BusinessLogicEditRoleUsingPUT(service.EditRoleUsingPUT))

	routes.FindAllFieldsUsingGET.RouterGroup = routes.Group("")
	routes.FindAllFieldsUsingGET.Post = routes.FindAllFieldsUsingGET.Group("")
	routes.FindAllFieldsUsingGET.Post.GET(ginizePath("/schemas/_fields"), schema_controller.BusinessLogicFindAllFieldsUsingGET(service.FindAllFieldsUsingGET))

	routes.FindEntityFieldsUsingGET.RouterGroup = routes.Group("")
	routes.FindEntityFieldsUsingGET.Post = routes.FindEntityFieldsUsingGET.Group("")
	routes.FindEntityFieldsUsingGET.Post.GET(ginizePath("/schemas/fields"), schema_controller.BusinessLogicFindEntityFieldsUsingGET(service.FindEntityFieldsUsingGET))

	routes.FindOneFieldUsingGET.RouterGroup = routes.Group("")
	routes.FindOneFieldUsingGET.Post = routes.FindOneFieldUsingGET.Group("")
	routes.FindOneFieldUsingGET.Post.GET(ginizePath("/schemas/_fields/{fid}"), schema_controller.BusinessLogicFindOneFieldUsingGET(service.FindOneFieldUsingGET))

	routes.FindOneUsingGET.RouterGroup = routes.Group("")
	routes.FindOneUsingGET.Post = routes.FindOneUsingGET.Group("")
	routes.FindOneUsingGET.Post.GET(ginizePath("/api/{entity}/{id}"), data_controller.BusinessLogicFindOneUsingGET(service.FindOneUsingGET))

	routes.FindOneUsingGET1.RouterGroup = routes.Group("")
	routes.FindOneUsingGET1.Post = routes.FindOneUsingGET1.Group("")
	routes.FindOneUsingGET1.Post.GET(ginizePath("/schemas/_entitys/{eid}"), schema_controller.BusinessLogicFindOneUsingGET1(service.FindOneUsingGET1))

	routes.FindRoleUsingGET.RouterGroup = routes.Group("")
	routes.FindRoleUsingGET.Post = routes.FindRoleUsingGET.Group("")
	routes.FindRoleUsingGET.Post.GET(ginizePath("/datasource/_datasource/{datasourceId}"), data_source_controller.BusinessLogicFindRoleUsingGET(service.FindRoleUsingGET))

	routes.FindRoleUsingGET1.RouterGroup = routes.Group("")
	routes.FindRoleUsingGET1.Post = routes.FindRoleUsingGET1.Group("")
	routes.FindRoleUsingGET1.Post.GET(ginizePath("/role/_roles/{roleId}"), role_controller.BusinessLogicFindRoleUsingGET1(service.FindRoleUsingGET1))

	routes.FindUserUsingGET.RouterGroup = routes.Group("")
	routes.FindUserUsingGET.Post = routes.FindUserUsingGET.Group("")
	routes.FindUserUsingGET.Post.GET(ginizePath("/permission/_permission/{id}"), permission_controller.BusinessLogicFindUserUsingGET(service.FindUserUsingGET))

	routes.FindUserUsingGET1.RouterGroup = routes.Group("")
	routes.FindUserUsingGET1.Post = routes.FindUserUsingGET1.Group("")
	routes.FindUserUsingGET1.Post.GET(ginizePath("/user/_users/{userId}"), user_controller.BusinessLogicFindUserUsingGET1(service.FindUserUsingGET1))

	routes.GetAuthenticatedUserUsingGET.RouterGroup = routes.Group("")
	routes.GetAuthenticatedUserUsingGET.Post = routes.GetAuthenticatedUserUsingGET.Group("")
	routes.GetAuthenticatedUserUsingGET.Post.GET(ginizePath("/user/me"), user_controller.BusinessLogicGetAuthenticatedUserUsingGET(service.GetAuthenticatedUserUsingGET))

	routes.GetSchemasUsingGET.RouterGroup = routes.Group("")
	routes.GetSchemasUsingGET.Post = routes.GetSchemasUsingGET.Group("")
	routes.GetSchemasUsingGET.Post.GET(ginizePath("/schemas/_entitys"), schema_controller.BusinessLogicGetSchemasUsingGET(service.GetSchemasUsingGET))

	routes.ListUsingGET.RouterGroup = routes.Group("")
	routes.ListUsingGET.Post = routes.ListUsingGET.Group("")
	routes.ListUsingGET.Post.GET(ginizePath("/datasource/_datasource"), data_source_controller.BusinessLogicListUsingGET(service.ListUsingGET))

	routes.ListUsingGET1.RouterGroup = routes.Group("")
	routes.ListUsingGET1.Post = routes.ListUsingGET1.Group("")
	routes.ListUsingGET1.Post.GET(ginizePath("/permission/_permission"), permission_controller.BusinessLogicListUsingGET1(service.ListUsingGET1))

	routes.ListUsingGET2.RouterGroup = routes.Group("")
	routes.ListUsingGET2.Post = routes.ListUsingGET2.Group("")
	routes.ListUsingGET2.Post.GET(ginizePath("/role/_roles"), role_controller.BusinessLogicListUsingGET2(service.ListUsingGET2))

	routes.ListUsingGET3.RouterGroup = routes.Group("")
	routes.ListUsingGET3.Post = routes.ListUsingGET3.Group("")
	routes.ListUsingGET3.Post.GET(ginizePath("/user/_users"), user_controller.BusinessLogicListUsingGET3(service.ListUsingGET3))

	routes.RefreshAndGetAuthenticationTokenUsingGET.RouterGroup = routes.Group("")
	routes.RefreshAndGetAuthenticationTokenUsingGET.Post = routes.RefreshAndGetAuthenticationTokenUsingGET.Group("")
	routes.RefreshAndGetAuthenticationTokenUsingGET.Post.GET(ginizePath("/refresh"), authentication_rest_controller.BusinessLogicRefreshAndGetAuthenticationTokenUsingGET(service.RefreshAndGetAuthenticationTokenUsingGET))

	routes.ResetCurrentDsUsingPUT.RouterGroup = routes.Group("")
	routes.ResetCurrentDsUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.ResetCurrentDsUsingPUT.Post = routes.ResetCurrentDsUsingPUT.Group("")
	routes.ResetCurrentDsUsingPUT.Post.PUT(ginizePath("/schemas/resetCurrentDs/{dataSourceId}"), schema_controller.BusinessLogicResetCurrentDsUsingPUT(service.ResetCurrentDsUsingPUT))

	routes.SyncSchemasUsingPUT.RouterGroup = routes.Group("")
	routes.SyncSchemasUsingPUT.RouterGroup.Use(middleware.ContentTypes("application/json"))
	routes.SyncSchemasUsingPUT.Post = routes.SyncSchemasUsingPUT.Group("")
	routes.SyncSchemasUsingPUT.Post.PUT(ginizePath("/schemas/sync/{dataSourceId}"), schema_controller.BusinessLogicSyncSchemasUsingPUT(service.SyncSchemasUsingPUT))

	return routes
}

// API defines the API service.
type API struct {
	Routes  *Routes
	config  *Config
	server  *http.Server
	Title   string
	Version string
}

// NewAPI initializes a new API service.
func NewAPI(svc Service, config *Config) *API {
	if !config.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	api := &API{
		Routes:  configureRoutes(svc, !config.AuthDisabled, config.TokenURL),
		config:  config,
		Title:   "DataHive RESTful APIs",
		Version: "1.0",
	}

	// enable pprof http endpoints in debug mode
	if config.Debug {
		pprof.Register(api.Routes.Engine, nil)
	}

	// set logrus logger to TextFormatter with no colors
	log.SetFormatter(&log.TextFormatter{DisableColors: true})

	api.server = &http.Server{
		Addr:         config.Address,
		Handler:      api.Routes.Engine,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if !config.WellKnownDisabled {
		api.Routes.configureWellKnown(svc.Healthy)
	}

	// configure healthz endpoint
	api.Routes.GET("/healthz", healthHandler(svc.Healthy))

	// configure info endpoint
	// api.Routes.GET("/info", infoHandler(svc.Info))

	return api
}

// Run runs the API server it will listen on either HTTP or HTTPS depending on
// the config passed to NewAPI.
func (a *API) Run() error {
	log.Infof("Serving '%s - %s' on address %s", a.Title, a.Version, a.server.Addr)
	if a.config.InsecureHTTP {
		return a.server.ListenAndServe()
	}
	return a.server.ListenAndServeTLS(a.config.TLSCertFile, a.config.TLSKeyFile)
}

// Shutdown will gracefully shutdown the API server.
func (a *API) Shutdown() error {
	return a.server.Shutdown(context.Background())
}

// RunWithSigHandler runs the API server with SIGTERM handling automatically
// enabled. The server will listen for a SIGTERM signal and gracefully shutdown
// the web server.
// It's possible to optionally pass any number shutdown functions which will
// execute one by one after the webserver has been shutdown successfully.
func (a *API) RunWithSigHandler(shutdown ...func() error) error {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM)

	go func() {
		<-sigCh
		a.Shutdown()
	}()

	err := a.Run()
	if err != nil {
		if err != http.ErrServerClosed {
			return err
		}
	}

	for _, fn := range shutdown {
		err := fn()
		if err != nil {
			return err
		}
	}

	return nil
}

// vim: ft=go
