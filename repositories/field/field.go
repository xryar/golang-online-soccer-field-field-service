package repositories

import (
	"context"
	"field-service/domain/dto"
	"field-service/domain/models"

	"gorm.io/gorm"
)

type FieldRepository struct {
	db *gorm.DB
}

type IFieldRepository interface {
	FindAllWithPagination(context.Context, *dto.FieldRequestParam) ([]models.Field, int64, error)
	FindAllWithoutPagination(context.Context) ([]models.Field, error)
	FindByUUID(context.Context, string) (*models.Field, error)
	Create(context.Context, *models.Field) (*models.Field, error)
	Update(context.Context, string, *models.Field) (*models.Field, error)
	Delete(context.Context, string) error
}

func NewFieldRepository(db *gorm.DB) IFieldRepository {
	return &FieldRepository{db: db}
}

func (fr *FieldRepository) FindAllWithPagination(context.Context, *dto.FieldRequestParam) ([]models.Field, int64, error)

func (fr *FieldRepository) FindAllWithoutPagination(context.Context) ([]models.Field, error)

func (fr *FieldRepository) FindByUUID(context.Context, string) (*models.Field, error)

func (fr *FieldRepository) Create(context.Context, *models.Field) (*models.Field, error)

func (fr *FieldRepository) Update(context.Context, string, *models.Field) (*models.Field, error)

func (fr *FieldRepository) Delete(context.Context, string) error
