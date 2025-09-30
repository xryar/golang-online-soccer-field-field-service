package services

import (
	"context"
	"field-service/common/util"
	"field-service/domain/dto"
	"field-service/repositories"
	"fmt"
	"time"
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
	fieldSchedules, total, err := fs.repository.GetFieldSchedule().FindAllWithPagination(ctx, param)
	if err != nil {
		return nil, err
	}

	fieldScheduleResults := make([]dto.FieldScheduleResponse, 0, len(fieldSchedules))
	for _, schedule := range fieldSchedules {
		fieldScheduleResults = append(fieldScheduleResults, dto.FieldScheduleResponse{
			UUID:         schedule.UUID,
			FieldName:    schedule.Field.Name,
			Date:         schedule.Date.Format("2006-01-02"),
			PricePerHour: schedule.Field.PricePerHour,
			Status:       schedule.Status.GetString(),
			Time:         fmt.Sprintf("%s - %s", schedule.Time.StartTme, schedule.Time.EndTime),
			CreatedAt:    schedule.CreatedAt,
			UpdatedAt:    schedule.UpdatedAt,
		})
	}

	pagination := &util.PaginationParam{
		Count: total,
		Limit: param.Limit,
		Page:  param.Page,
		Data:  fieldScheduleResults,
	}

	response := util.GeneratePagination(*pagination)
	return &response, nil
}

func (fs *FieldScheduleService) GetAllByFieldIDAndDate(ctx context.Context, uuid, date string) ([]dto.FieldScheduleForBookingResponse, error) {
	field, err := fs.repository.GetFieldSchedule().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	fieldSchedules, err := fs.repository.GetFieldSchedule().FindAllByFieldIDAndDate(ctx, int(field.ID), date)
	if err != nil {
		return nil, err
	}

	fieldScheduleResults := make([]dto.FieldScheduleForBookingResponse, 0, len(fieldSchedules))
	for _, schedule := range fieldSchedules {
		pricePerHour := float64(schedule.Field.PricePerHour)
		fieldScheduleResults = append(fieldScheduleResults, dto.FieldScheduleForBookingResponse{
			UUID:         schedule.UUID,
			Date:         fs.convertMonthName(schedule.Date.Format(time.DateOnly)),
			Time:         fmt.Sprintf("%s - %s", schedule.Time.StartTme, schedule.Time.EndTime),
			Status:       schedule.Status.GetString(),
			PricePerHour: util.RupiahFormat(&pricePerHour),
		})
	}

	return fieldScheduleResults, nil
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

func (fs *FieldScheduleService) convertMonthName(inputDate string) string {
	date, err := time.Parse(time.DateOnly, inputDate)
	if err != nil {
		return ""
	}

	indonesiaMonth := map[string]string{
		"Jan": "Jan",
		"Feb": "Feb",
		"Mar": "Mar",
		"Apr": "Apr",
		"May": "Mei",
		"Jun": "Jun",
		"Jul": "Jul",
		"Aug": "Agu",
		"Sep": "Sep",
		"Oct": "Okt",
		"Nov": "Nov",
		"Dec": "Des",
	}

	formattedDate := date.Format("02 Jan")
	day := formattedDate[:3]
	month := formattedDate[3:]
	formattedDate = fmt.Sprintf("%s %s", day, indonesiaMonth[month])

	return formattedDate
}
