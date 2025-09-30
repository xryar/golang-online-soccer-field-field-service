package services

import (
	"context"
	"field-service/domain/dto"
	"field-service/repositories"
)

type TimeService struct {
	repository repositories.IRegistryRepository
}

type ITimeService interface {
	GetAll(context.Context) ([]dto.TimeResponse, error)
	GetByUUID(context.Context, string) (*dto.TimeResponse, error)
	Create(context.Context, *dto.TimeRequest) (*dto.TimeResponse, error)
}

func NewTimeService(repository repositories.IRegistryRepository) ITimeService {
	return &TimeService{repository: repository}
}

func (ts *TimeService) GetAll(ctx context.Context) ([]dto.TimeResponse, error) {}

func (ts *TimeService) GetByUUID(ctx context.Context, uuid string) (*dto.TimeResponse, error) {}

func (ts *TimeService) Create(ctx context.Context, req *dto.TimeRequest) (*dto.TimeResponse, error) {}
