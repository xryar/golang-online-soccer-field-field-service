package routes

import (
	"field-service/clients"
	"field-service/constants"
	"field-service/controllers"
	"field-service/middlewares"

	"github.com/gin-gonic/gin"
)

type FieldScheduleRoute struct {
	controller controllers.IRegistryController
	group      *gin.RouterGroup
	client     clients.IRegistryClient
}

type IFieldSchedule interface {
	Run()
}

func NewFieldScheduleRoute(router *gin.Engine, controller controllers.IRegistryController, client clients.IRegistryClient) IFieldSchedule {
	return &FieldScheduleRoute{
		controller: controller,
		group:      router.Group("/field"),
		client:     client,
	}
}

func (fr *FieldScheduleRoute) Run() {
	group := fr.group.Group("/field/schedule")
	group.GET("", middlewares.AuthenticateWithoutToken(), fr.controller.GetFieldSchedule().GetAllByFieldIDAndDate)
	group.PATCH("", middlewares.AuthenticateWithoutToken(), fr.controller.GetFieldSchedule().UpdateStatus)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, fr.client), fr.controller.GetFieldSchedule().GetAllWithPagination)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, fr.client), fr.controller.GetFieldSchedule().GetByUUID)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetFieldSchedule().Create)
	group.POST("/one-month", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetFieldSchedule().GenerateScheduleForOneMonth)
	group.PUT("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetFieldSchedule().Update)
	group.DELETE("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, fr.client), fr.controller.GetFieldSchedule().Delete)
}
