package routes

import (
	"field-service/clients"
	"field-service/controllers"
	fieldRoutes "field-service/routes/field"
	fieldScheduleRoutes "field-service/routes/fieldSchedule"
	timeRoutes "field-service/routes/time"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	controller controllers.IRegistryController
	group      *gin.RouterGroup
	client     clients.IRegistryClient
}

type IRegistry interface {
	Serve()
}

func NewRegistryRoute(controller controllers.IRegistryController, group *gin.RouterGroup, client clients.IRegistryClient) IRegistry {
	return &Registry{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (r *Registry) fieldRoute() fieldRoutes.IFieldRoute {
	return fieldRoutes.NewFieldRoute(r.controller, r.group, r.client)
}

func (r *Registry) fieldScheduleRoute() fieldScheduleRoutes.IFieldScheduleRoute {
	return fieldScheduleRoutes.NewFieldScheduleRoute(r.controller, r.group, r.client)
}

func (r *Registry) timeRoute() timeRoutes.ITimeRoute {
	return timeRoutes.NewTimeRoute(r.controller, r.group, r.client)
}

func (r *Registry) Serve() {
	r.fieldRoute().Run()
	r.fieldScheduleRoute().Run()
	r.timeRoute().Run()
}
