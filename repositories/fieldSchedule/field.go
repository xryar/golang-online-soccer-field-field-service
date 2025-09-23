package repositories

import (
	"context"
	"errors"
	errWrap "field-service/common/error"
	"field-service/constants"
	errConstant "field-service/constants/error"
	errFieldSchedule "field-service/constants/error/fieldSchedule"
	"field-service/domain/dto"
	"field-service/domain/models"
	"fmt"

	"gorm.io/gorm"
)

type FieldScheduleRepository struct {
	db *gorm.DB
}

type IFieldScheduleRepository interface {
	FindAllWithPagination(context.Context, *dto.FieldScheduleRequestParam) ([]models.FieldSchedule, int64, error)
	FindAllByFieldIDAndDate(context.Context, int, string) ([]models.FieldSchedule, error)
	FindByUUID(context.Context, string) (*models.FieldSchedule, error)
	FindByDateAndTimeID(context.Context, string, int, int) (*models.FieldSchedule, error)
	Create(context.Context, []models.FieldSchedule) error
	Update(context.Context, string, *models.FieldSchedule) (*models.FieldSchedule, error)
	UpdateStatus(context.Context, constants.FieldScheduleStatus) (*models.FieldSchedule, error)
	Delete(context.Context, string) error
}

func NewFieldScheduleRepository(db *gorm.DB) IFieldScheduleRepository {
	return &FieldScheduleRepository{db: db}
}

func (fr *FieldScheduleRepository) FindAllWithPagination(ctx context.Context, param *dto.FieldScheduleRequestParam) ([]models.FieldSchedule, int64, error) {
	var (
		fieldSchedules []models.FieldSchedule
		sort           string
		total          int64
	)
	if param.SortColumn != nil {
		sort = fmt.Sprintf("%s %s", *param.SortColumn, *param.SortOrder)
	} else {
		sort = "created_at desc"
	}

	limit := param.Limit
	offset := (param.Page - 1) * limit
	err := fr.db.WithContext(ctx).Preload("Field").Preload("Time").Limit(limit).Offset(offset).Order(sort).Find(&fieldSchedules).Error
	if err != nil {
		return nil, 0, errWrap.WrapError(errConstant.ErrSQLError)
	}

	err = fr.db.WithContext(ctx).Model(&fieldSchedules).Count(&total).Error
	if err != nil {
		return nil, 0, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return fieldSchedules, total, nil

}

func (fr *FieldScheduleRepository) FindAllByFieldIDAndDate(ctx context.Context, fieldId int, date string) ([]models.FieldSchedule, error) {
	var fieldSchedules []models.FieldSchedule
	err := fr.db.
		WithContext(ctx).
		Preload("Field").
		Preload("Time").
		Where("field_id ?", fieldId).
		Where("date = ?", date).
		Joins("LEFT JOIN times ON field_schedules.time_id = time.id").
		Order("times.start_time asc").
		Find(&fieldSchedules).
		Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return fieldSchedules, nil
}

func (fr *FieldScheduleRepository) FindByUUID(ctx context.Context, uuid string) (*models.FieldSchedule, error) {
	var fieldSchedule models.FieldSchedule
	err := fr.db.WithContext(ctx).Preload("Field").Preload("Time").Where("uuid = ?", uuid).First(&fieldSchedule).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errFieldSchedule.ErrFieldScheduleNotFound)
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &fieldSchedule, nil
}

func (fr *FieldScheduleRepository) FindByDateAndTimeID(ctx context.Context, date string, timeID int, fieldID int) (*models.FieldSchedule, error) {
	var fieldSchedule models.FieldSchedule
	err := fr.db.
		WithContext(ctx).
		Where("date = ?", date).
		Where("time_id = ?", timeID).
		Where("field_id = ?", fieldID).
		First(&fieldSchedule).
		Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errFieldSchedule.ErrFieldScheduleNotFound)
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &fieldSchedule, nil
}

func (fr *FieldScheduleRepository) Create(ctx context.Context, req []models.FieldSchedule) error {
	err := fr.db.WithContext(ctx).Create(&req).Error
	if err != nil {
		return errWrap.WrapError(errConstant.ErrSQLError)
	}

	return nil
}

func (fr *FieldScheduleRepository) Update(ctx context.Context, uuid string, req *models.Field) (*models.Field, error) {
	field := models.Field{
		Code:         req.Code,
		Name:         req.Name,
		Images:       req.Images,
		PricePerHour: req.PricePerHour,
	}

	err := fr.db.WithContext(ctx).Where("uuid = ?", uuid).Updates(&field).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &field, nil
}

func (fr *FieldScheduleRepository) Delete(ctx context.Context, uuid string) error {
	err := fr.db.WithContext(ctx).Where("uuid = ?", uuid).Delete(&models.Field{}).Error
	if err != nil {
		return errWrap.WrapError(errConstant.ErrSQLError)
	}

	return nil
}
