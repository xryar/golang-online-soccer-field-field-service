package controllers

import (
	fieldController "field-service/controllers/field"
	fieldScheduleController "field-service/controllers/fieldSchedule"
	timeController "field-service/controllers/time"
	"field-service/services"
)

type Registry struct {
	service services.IRegistryService
}

type IRegistryController interface {
	GetField() fieldController.IFieldController
	GetFieldSchedule() fieldScheduleController.IFieldScheduleController
	GetTime() timeController.ITimeController
}

func NewRegistryController(service services.IRegistryService) IRegistryController {
	return &Registry{service: service}
}

func (r *Registry) GetField() fieldController.IFieldController {
	return fieldController.NewFieldController(r.service)
}

func (r *Registry) GetFieldSchedule() fieldScheduleController.IFieldScheduleController {
	return fieldScheduleController.NewFieldScheduleController(r.service)
}

func (r *Registry) GetTime() timeController.ITimeController {
	return timeController.NewTimeController(r.service)
}
