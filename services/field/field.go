package services

import (
	"context"
	"field-service/common/gcs"
	"field-service/common/util"
	"field-service/domain/dto"
	"field-service/repositories"
)

type FieldService struct {
	repository repositories.IRegistryRepository
	gcs        gcs.IGCSClient
}

type IFieldService interface {
	GetAllPagination(context.Context, *dto.FieldRequestParam) (*util.PaginationResult, error)
	GetAllWithoutPagination(context.Context) ([]dto.FieldResponse, error)
	GetByUUID(context.Context, string) (*dto.FieldResponse, error)
	Create(context.Context, *dto.FieldRequest) (*dto.FieldResponse, error)
	Update(context.Context, string, *dto.UpdateFieldRequest) (*dto.FieldResponse, error)
	Delete(context.Context, string) error
}

func NewFieldService(repository repositories.IRegistryRepository, gcs gcs.IGCSClient) IFieldService {
	return &FieldService{
		repository: repository,
		gcs:        gcs,
	}
}

func (fs *FieldService) GetAllPagination(ctx context.Context, param *dto.FieldRequestParam) (*util.PaginationResult, error) {
}

func (fs *FieldService) GetAllWithoutPagination(ctx context.Context) ([]dto.FieldResponse, error) {

}

func (fs *FieldService) GetByUUID(ctx context.Context, fieldUUID string) (*dto.FieldResponse, error) {
}

func (fs *FieldService) Create(ctx context.Context, field *dto.FieldRequest) (*dto.FieldResponse, error) {
}

func (fs *FieldService) Update(ctx context.Context, fieldUUID string, field *dto.UpdateFieldRequest) (*dto.FieldResponse, error) {
}

func (fs *FieldService) Delete(ctx context.Context, fieldUUID string) error {}
