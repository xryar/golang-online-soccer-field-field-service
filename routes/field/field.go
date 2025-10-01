package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"

	"github.com/gin-gonic/gin"
)

type FieldRoute struct {
	controller controllers.IRegistryController
	group      *gin.RouterGroup
	client     clients.IRegistryClient
}

type IFieldRoute interface {
	Run()
}

func NewFieldRoute(router *gin.Engine, controller controllers.IRegistryController, client clients.IRegistryClient) IFieldRoute {
	return &FieldRoute{
		controller: controller,
		group:      router.Group("/field"),
		client:     client,
	}
}

func (fr *FieldRoute) Run() {
	group := fr.group.Group("/field")
	group.GET("", middlewares.AuthenticateWithoutToken(), fr.controller.GetField().GetAllWithoutPagination)
	group.GET("/:uuid", middlewares.AuthenticateWithoutToken(), fr.controller.GetField().GetByUUID)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, fr.client), fr.controller.GetField().GetAllWithPagination)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetField().Create)
	group.PUT("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetField().Update)
	group.DELETE("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetField().Delete)
}
