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

type ITimeRoute interface {
	Run()
}

func NewTimeRoute(controller controllers.IRegistryController, group *gin.RouterGroup, client clients.IRegistryClient) ITimeRoute {
	return &TimeRoute{
		controller: controller,
		group:      group,
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
