package repositories

import (
	fieldRepo "field-service/repositories/field"
	fieldScheduleRepo "field-service/repositories/fieldSchedule"
	timeRepo "field-service/repositories/time"

	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRegistryRepository interface {
	GetField() fieldRepo.IFieldRepository
	GetFieldSchedule() fieldScheduleRepo.IFieldScheduleRepository
	GetTime() timeRepo.ITimeRepository
}

func NewRegistryRepository(db *gorm.DB) IRegistryRepository {
	return &Registry{db: db}
}

func (r *Registry) GetField() fieldRepo.IFieldRepository {
	return fieldRepo.NewFieldRepository(r.db)
}

func (r *Registry) GetFieldSchedule() fieldScheduleRepo.IFieldScheduleRepository {
	return fieldScheduleRepo.NewFieldScheduleRepository(r.db)
}

func (r *Registry) GetTime() timeRepo.ITimeRepository {
	return timeRepo.NewTimeRepository(r.db)
}
