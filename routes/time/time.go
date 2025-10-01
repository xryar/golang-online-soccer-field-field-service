package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"

	"github.com/gin-gonic/gin"
)

type TimeRoute struct {
	controller controllers.IRegistryController
	group      *gin.RouterGroup
	client     clients.IRegistryClient
}

type ITime interface {
	Run()
}

func NewTimeRoute(router *gin.Engine, controller controllers.IRegistryController, client clients.IRegistryClient) ITime {
	return &TimeRoute{
		controller: controller,
		group:      router.Group("/time"),
		client:     client,
	}
}

func (fr *TimeRoute) Run() {
	group := fr.group.Group("/time")
	group.Use(middlewares.Authenticate())
	group.GET("", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetTime().GetAll)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetTime().GetByUUID)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetTime().Create)
}
