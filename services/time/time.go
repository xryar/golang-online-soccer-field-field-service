package services

import (
	"context"
	"field-service/domain/dto"
	"field-service/domain/models"
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

func (ts *TimeService) GetAll(ctx context.Context) ([]dto.TimeResponse, error) {
	times, err := ts.repository.GetTime().FindAll(ctx)
	if err != nil {
		return nil, err
	}

	timeResults := make([]dto.TimeResponse, 0, len(times))
	for _, time := range times {
		timeResults = append(timeResults, dto.TimeResponse{
			UUID:      time.UUID,
			StartTime: time.StartTime,
			EndTime:   time.EndTime,
			CreatedAt: time.CreatedAt,
			UpdatedAt: time.UpdatedAt,
		})
	}

	return timeResults, nil
}

func (ts *TimeService) GetByUUID(ctx context.Context, uuid string) (*dto.TimeResponse, error) {
	time, err := ts.repository.GetTime().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	timeResult := dto.TimeResponse{
		UUID:      time.UUID,
		StartTime: time.StartTime,
		EndTime:   time.EndTime,
		CreatedAt: time.CreatedAt,
		UpdatedAt: time.UpdatedAt,
	}

	return &timeResult, nil
}

func (ts *TimeService) Create(ctx context.Context, req *dto.TimeRequest) (*dto.TimeResponse, error) {
	time := &dto.TimeRequest{
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}

	timeResult, err := ts.repository.GetTime().Create(ctx, &models.Time{
		StartTime: time.StartTime,
		EndTime:   time.EndTime,
	})
	if err != nil {
		return nil, err
	}

	response := dto.TimeResponse{
		UUID:      timeResult.UUID,
		StartTime: timeResult.StartTime,
		EndTime:   timeResult.EndTime,
		CreatedAt: timeResult.CreatedAt,
		UpdatedAt: timeResult.UpdatedAt,
	}

	return &response, nil
}
