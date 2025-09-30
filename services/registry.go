package services

import (
	"field-service/common/gcs"
	"field-service/repositories"
	fieldService "field-service/services/field"
	fieldScheduleService "field-service/services/fieldSchedule"
	timeService "field-service/services/time"
)

type Registry struct {
	repository repositories.IRegistryRepository
	gcs        gcs.IGCSClient
}

type IRegistryService interface {
	GetField() fieldService.IFieldService
	GetFieldSchedule() fieldScheduleService.IFieldScheduleService
	GetTime() timeService.ITimeService
}

func NewRegistryService(repository repositories.IRegistryRepository, gcs gcs.IGCSClient) IRegistryService {
	return &Registry{
		repository: repository,
		gcs:        gcs,
	}
}

func (r *Registry) GetField() fieldService.IFieldService {
	return fieldService.NewFieldService(r.repository, r.gcs)
}

func (r *Registry) GetFieldSchedule() fieldScheduleService.IFieldScheduleService {
	return fieldScheduleService.NewFieldScheduleService(r.repository)
}

func (r *Registry) GetTime() timeService.ITimeService {
	return timeService.NewTimeService(r.repository)
}
