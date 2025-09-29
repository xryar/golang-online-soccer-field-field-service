package repositories

import (
	"context"
	"errors"
	errWrap "field-service/common/error"
	errConstant "field-service/constants/error"
	errTime "field-service/constants/error/time"
	"field-service/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TimeRepository struct {
	db *gorm.DB
}

type ITimeRepository interface {
	FindAll(context.Context) ([]models.Time, error)
	FindByUUID(context.Context, string) (*models.Time, error)
	FindByID(context.Context, int) (*models.Time, error)
	Create(context.Context, *models.Time) (*models.Time, error)
}

func NewTimeRepository(db *gorm.DB) ITimeRepository {
	return &TimeRepository{db: db}
}

func (tr *TimeRepository) FindAll(ctx context.Context) ([]models.Time, error) {
	var times []models.Time
	err := tr.db.WithContext(ctx).Find(&times).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return times, nil
}

func (tr *TimeRepository) FindByUUID(ctx context.Context, uuid string) (*models.Time, error) {
	var time models.Time
	err := tr.db.WithContext(ctx).Where("uuid = ?", uuid).First(&time).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errTime.ErrTimeNotFound)
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &time, nil

}

func (tr *TimeRepository) FindByID(ctx context.Context, id int) (*models.Time, error) {
	var time models.Time
	err := tr.db.WithContext(ctx).Where("id = ?", id).First(&time).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errWrap.WrapError(errTime.ErrTimeNotFound)
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &time, nil
}

func (tr *TimeRepository) Create(ctx context.Context, time *models.Time) (*models.Time, error) {
	time.UUID = uuid.New()
	err := tr.db.WithContext(ctx).Create(&time).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return time, nil
}
