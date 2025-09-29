package services

import (
	"context"
	"field-service/common/util"
	"field-service/domain/dto"
	"field-service/repositories"
)

type FieldScheduleService struct {
	repository repositories.IRegistryRepository
}

type IFieldScheduleService interface {
	GetAllWithPagination(context.Context, *dto.FieldScheduleRequestParam) (*util.PaginationResult, error)
	GetAllByFieldIDAndDate(context.Context, string, string) ([]dto.FieldScheduleForBookingResponse, error)
	GetByUUID(context.Context, string) (*dto.FieldScheduleResponse, error)
	GenerateScheduleForOneMonth(context.Context, *dto.GenerateFieldScheduleForOneMonthRequest) error
	Create(context.Context, *dto.FieldScheduleRequest) error
	Update(context.Context, *dto.FieldScheduleRequest) (*dto.FieldScheduleResponse, error)
	UpdateStatus(context.Context, *dto.UpdateStatusFieldScheduleRequest) error
	Delete(context.Context, string) error
}

func NewFieldScheduleService(repository repositories.IRegistryRepository) IFieldScheduleService {
	return &FieldScheduleService{repository: repository}
}

func (fs *FieldScheduleService) GetAllWithPagination(ctx context.Context, param *dto.FieldScheduleRequestParam) (*util.PaginationResult, error) {
}

func (fs *FieldScheduleService) GetAllByFieldIDAndDate(ctx context.Context, fieldID string, date string) ([]dto.FieldScheduleForBookingResponse, error) {
}

func (fs *FieldScheduleService) GetByUUID(ctx context.Context, uuid string) (*dto.FieldScheduleResponse, error) {
}

func (fs *FieldScheduleService) GenerateScheduleForOneMonth(ctx context.Context, req *dto.GenerateFieldScheduleForOneMonthRequest) error {
}

func (fs *FieldScheduleService) Create(ctx context.Context, req *dto.FieldScheduleRequest) error {}

func (fs *FieldScheduleService) Update(ctx context.Context, req *dto.FieldScheduleRequest) (*dto.FieldScheduleResponse, error) {
}

func (fs *FieldScheduleService) UpdateStatus(ctx context.Context, req *dto.UpdateStatusFieldScheduleRequest) error {
}

func (fs *FieldScheduleService) Delete(ctx context.Context, uuid string) error {}
