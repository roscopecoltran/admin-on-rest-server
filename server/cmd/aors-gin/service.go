package main

import (
	"net/http"

	"github.com/roscopecoltran/admin-on-rest-server/server/api"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/apply_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/authentication_rest_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/data_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/data_source_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/permission_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/role_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/schema_controller"
	"github.com/roscopecoltran/admin-on-rest-server/server/restapi/operations/user_controller"

	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp"
	// "github.com/roscopecoltran/admin-on-rest-server/server/database"
	"github.com/roscopecoltran/admin-on-rest-server/server/models"

	"github.com/gin-gonic/gin"
)

type AdminOnRestServer struct {
	Health bool
}

func (s *AdminOnRestServer) Healthy() bool {
	return s.Health
}

func (s *AdminOnRestServer) Info(ctx *gin.Context) *api.Response {
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) AddDataSourceUsingPOST(ctx *gin.Context, params *data_source_controller.AddDataSourceUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	response := make(map[string]string)
	var dataSource models.DataSource
	var queryRes models.DataSource
	db := ctx.MustGet("db").(*gorm.DB)
	db.Where("cluster_name = ? && db_name = ?", params.DataSource.ClusterName, params.DataSource.DbName).First(&queryRes)
	if queryRes.DbName != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&dataSource)
		response["status"] = "success"
	}
	return &api.Response{Code: http.StatusOK, Body: response}
}

func (s *AdminOnRestServer) AddEntityUsingPOST(ctx *gin.Context, params *schema_controller.AddEntityUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) AddFieldUsingPOST(ctx *gin.Context, params *schema_controller.AddFieldUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) AddPermissionUsingPOST(ctx *gin.Context, params *permission_controller.AddPermissionUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) AddRoleUsingPOST(ctx *gin.Context, params *role_controller.AddRoleUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) AddUserUsingPOST(ctx *gin.Context, params *user_controller.AddUserUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	response := make(map[string]string)
	db := ctx.MustGet("db").(*gorm.DB)
	var queryRes models.User
	db.Where("email = ?", params.User.Email).First(&queryRes)
	if queryRes.Email != "" {
		response["msg"] = "Duplicate resource."
		response["status"] = "error"
	} else {
		db.Create(&params.User)
		response["msg"] = "New user created."
		response["status"] = "success"
	}
	return &api.Response{Code: http.StatusOK, Body: response}
}

func (s *AdminOnRestServer) ApplyUsingPOST(ctx *gin.Context, params *apply_controller.ApplyUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	db := ctx.MustGet("db").(*gorm.DB)
	var queryRes models.Apply
	db.Where("email = ?", params.Apply.Email).First(&queryRes)
	response := make(map[string]string)
	if queryRes.Email != "" {
		response["status"] = "error"
		response["msg"] = "Duplicate resource."
	} else {
		db.Create(&params.Apply)
		response["msg"] = "Success, new user created."
		response["status"] = "success"
	}
	return &api.Response{Code: http.StatusOK, Body: response}
}

func (s *AdminOnRestServer) CreateAuthenticationTokenUsingPOST(ctx *gin.Context, params *authentication_rest_controller.CreateAuthenticationTokenUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) DataMutationUsingDELETE(ctx *gin.Context, params *data_controller.DataMutationUsingDELETEParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) DataMutationUsingPOST(ctx *gin.Context, params *data_controller.DataMutationUsingPOSTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) DataMutationUsingPUT(ctx *gin.Context, params *data_controller.DataMutationUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) DataQueryUsingGET(ctx *gin.Context, params *data_controller.DataQueryUsingGETParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) EditDataSourceUsingPUT(ctx *gin.Context, params *data_source_controller.EditDataSourceUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) EditEntityUsingPUT(ctx *gin.Context, params *schema_controller.EditEntityUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) EditFieldUsingPUT(ctx *gin.Context, params *schema_controller.EditFieldUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) EditFieldUsingPUT1(ctx *gin.Context, params *user_controller.EditFieldUsingPUT1Params) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) EditPermissionUsingPUT(ctx *gin.Context, params *permission_controller.EditPermissionUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) EditRoleUsingPUT(ctx *gin.Context, params *role_controller.EditRoleUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindAllFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindAllFieldsUsingGETParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindEntityFieldsUsingGET(ctx *gin.Context, params *schema_controller.FindEntityFieldsUsingGETParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindOneFieldUsingGET(ctx *gin.Context, params *schema_controller.FindOneFieldUsingGETParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindOneUsingGET(ctx *gin.Context, params *data_controller.FindOneUsingGETParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindOneUsingGET1(ctx *gin.Context, params *schema_controller.FindOneUsingGET1Params) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindRoleUsingGET(ctx *gin.Context, params *data_source_controller.FindRoleUsingGETParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindRoleUsingGET1(ctx *gin.Context, params *role_controller.FindRoleUsingGET1Params) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindUserUsingGET(ctx *gin.Context, params *permission_controller.FindUserUsingGETParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) FindUserUsingGET1(ctx *gin.Context, params *user_controller.FindUserUsingGET1Params) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) GetAuthenticatedUserUsingGET(ctx *gin.Context) *api.Response {
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) GetSchemasUsingGET(ctx *gin.Context) *api.Response {
	db := ctx.MustGet("db").(*gorm.DB)
	var queryRes []models.Entity
	response := make(map[string]interface{})
	if err := db.Select("*").Find(&queryRes).Error; err != nil {
		response["error"] = err.Error()
		return &api.Response{Code: 400, Body: response}
	}
	response["results"] = queryRes
	response["status"] = "success"
	return &api.Response{Code: http.StatusOK, Body: response}
}

func (s *AdminOnRestServer) ListUsingGET(ctx *gin.Context) *api.Response {
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) ListUsingGET1(ctx *gin.Context, params *permission_controller.ListUsingGET1Params) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) ListUsingGET2(ctx *gin.Context) *api.Response {
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) ListUsingGET3(ctx *gin.Context) *api.Response {
	db := ctx.MustGet("db").(*gorm.DB)
	var queryRes []models.User
	response := make(map[string]interface{})
	if err := db.Select("*").Find(&queryRes).Error; err != nil {
		response["error"] = err.Error()
		return &api.Response{Code: 400, Body: response}
	}
	response["results"] = queryRes
	response["status"] = "success"
	return &api.Response{Code: http.StatusOK, Body: response}
}

func (s *AdminOnRestServer) RefreshAndGetAuthenticationTokenUsingGET(ctx *gin.Context) *api.Response {
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) ResetCurrentDsUsingPUT(ctx *gin.Context, params *schema_controller.ResetCurrentDsUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}

func (s *AdminOnRestServer) SyncSchemasUsingPUT(ctx *gin.Context, params *schema_controller.SyncSchemasUsingPUTParams) *api.Response {
	pp.Println("params: ", params)
	// response := make(map[string]interface{})
	// db := ctx.MustGet("db").(*gorm.DB)
	// return &api.Response{Code: http.StatusOK, Body: response}
	return &api.Response{Code: http.StatusNotImplemented, Body: "Not Implemented"}
}
