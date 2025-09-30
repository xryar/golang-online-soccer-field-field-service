package services

import (
	"context"
	"field-service/common/util"
	"field-service/constants"
	errFieldSchedule "field-service/constants/error/fieldSchedule"
	"field-service/domain/dto"
	"field-service/domain/models"
	"field-service/repositories"
	"fmt"
	"time"

	"github.com/google/uuid"
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
	Update(context.Context, string, *dto.UpdateFieldScheduleRequets) (*dto.FieldScheduleResponse, error)
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
	fieldSchedule, err := fs.repository.GetFieldSchedule().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	response := &dto.FieldScheduleResponse{
		UUID:         fieldSchedule.UUID,
		FieldName:    fieldSchedule.Field.Name,
		PricePerHour: fieldSchedule.Field.PricePerHour,
		Date:         fieldSchedule.Date.Format(time.DateOnly),
		Status:       fieldSchedule.Status.GetString(),
		CreatedAt:    fieldSchedule.CreatedAt,
		UpdatedAt:    fieldSchedule.UpdatedAt,
	}

	return response, nil
}

func (fs *FieldScheduleService) Create(ctx context.Context, req *dto.FieldScheduleRequest) error {
	field, err := fs.repository.GetFieldSchedule().FindByUUID(ctx, req.FieldID)
	if err != nil {
		return err
	}

	fieldSchedules := make([]models.FieldSchedule, 0, len(req.TimeIDs))
	dateParsed, _ := time.Parse(time.DateOnly, req.Date)
	for _, timeID := range req.TimeIDs {
		scheduleTime, err := fs.repository.GetTime().FindByUUID(ctx, timeID)
		if err != nil {
			return err
		}

		schedule, err := fs.repository.GetFieldSchedule().FindByDateAndTimeID(ctx, req.Date, int(scheduleTime.ID), int(field.ID))
		if err != nil {
			return err
		}

		if schedule != nil {
			return errFieldSchedule.ErrFieldScheduleIsExist
		}

		fieldSchedules = append(fieldSchedules, models.FieldSchedule{
			UUID:    uuid.New(),
			FieldID: field.ID,
			TimeID:  scheduleTime.ID,
			Date:    dateParsed,
			Status:  constants.Available,
		})
	}

	err = fs.repository.GetFieldSchedule().Create(ctx, fieldSchedules)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FieldScheduleService) GenerateScheduleForOneMonth(ctx context.Context, req *dto.GenerateFieldScheduleForOneMonthRequest) error {
	field, err := fs.repository.GetFieldSchedule().FindByUUID(ctx, req.FieldID)
	if err != nil {
		return err
	}

	times, err := fs.repository.GetTime().FindAll(ctx)
	if err != nil {
		return err
	}

	numberOfDays := 30
	fieldSchedules := make([]models.FieldSchedule, 0, numberOfDays)
	now := time.Now().Add(time.Duration(1) * 24 * time.Hour)
	for i := 0; i < numberOfDays; i++ {
		currentDate := now.AddDate(0, 0, i)
		for _, item := range times {
			schedule, err := fs.repository.GetFieldSchedule().FindByDateAndTimeID(
				ctx,
				currentDate.Format(time.DateOnly),
				int(item.ID),
				int(field.ID),
			)
			if err != nil {
				return err
			}

			if schedule != nil {
				return errFieldSchedule.ErrFieldScheduleIsExist
			}

			fieldSchedules = append(fieldSchedules, models.FieldSchedule{
				UUID:    uuid.New(),
				FieldID: field.ID,
				TimeID:  item.ID,
				Date:    currentDate,
				Status:  constants.Available,
			})
		}
	}

	err = fs.repository.GetFieldSchedule().Create(ctx, fieldSchedules)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FieldScheduleService) Update(ctx context.Context, uuid string, req *dto.UpdateFieldScheduleRequets) (*dto.FieldScheduleResponse, error) {
	fieldSchedule, err := fs.repository.GetFieldSchedule().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	scheduleTime, err := fs.repository.GetTime().FindByUUID(ctx, req.TimeID)
	if err != nil {
		return nil, err
	}

	isTimeExist, err := fs.repository.GetFieldSchedule().FindByDateAndTimeID(ctx, req.Date, int(scheduleTime.ID), int(fieldSchedule.FieldID))
	if err != nil {
		return nil, err
	}

	if isTimeExist != nil && req.Date != fieldSchedule.Date.Format(time.DateOnly) {
		checkDate, err := fs.repository.GetFieldSchedule().FindByDateAndTimeID(
			ctx,
			req.Date,
			int(scheduleTime.ID),
			int(fieldSchedule.FieldID),
		)
		if err != nil {
			return nil, err
		}

		if checkDate != nil {
			return nil, errFieldSchedule.ErrFieldScheduleIsExist
		}
	}

	dateParsed, _ := time.Parse(time.DateOnly, req.Date)
	fieldResult, err := fs.repository.GetFieldSchedule().Update(ctx, uuid, &models.FieldSchedule{
		Date:   dateParsed,
		TimeID: scheduleTime.ID,
	})
	if err != nil {
		return nil, err
	}

	response := dto.FieldScheduleResponse{
		UUID:         fieldResult.UUID,
		FieldName:    fieldResult.Field.Name,
		Date:         fieldResult.Date.Format(time.DateOnly),
		PricePerHour: fieldResult.Field.PricePerHour,
		Status:       fieldResult.Status.GetString(),
		Time:         fmt.Sprintf("%s - %s", fieldResult.Time.StartTme, fieldResult.Time.EndTime),
	}

	return &response, nil
}

func (fs *FieldScheduleService) UpdateStatus(ctx context.Context, req *dto.UpdateStatusFieldScheduleRequest) error {
	for _, item := range req.FieldScheduleIDs {
		_, err := fs.repository.GetFieldSchedule().FindByUUID(ctx, item)
		if err != nil {
			return err
		}

		err = fs.repository.GetFieldSchedule().UpdateStatus(ctx, constants.Booked, item)
		if err != nil {
			return err
		}
	}

	return nil
}

func (fs *FieldScheduleService) Delete(ctx context.Context, uuid string) error {
	_, err := fs.repository.GetFieldSchedule().FindByUUID(ctx, uuid)
	if err != nil {
		return err
	}

	err = fs.repository.GetFieldSchedule().Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

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
