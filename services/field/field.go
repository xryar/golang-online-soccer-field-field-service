package services

import (
	"bytes"
	"context"
	"field-service/common/gcs"
	"field-service/common/util"
	errConstant "field-service/constants/error"
	"field-service/domain/dto"
	"field-service/domain/models"
	"field-service/repositories"
	"fmt"
	"io"
	"mime/multipart"
	"path"
	"time"
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
	fields, total, err := fs.repository.GetField().FindAllWithPagination(ctx, param)
	if err != nil {
		return nil, err
	}

	fieldResults := make([]dto.FieldResponse, 0, len(fields))
	for _, field := range fields {
		fieldResults = append(fieldResults, dto.FieldResponse{
			UUID:         field.UUID,
			Code:         field.Code,
			Name:         field.Name,
			PricePerHour: field.PricePerHour,
			Images:       field.Images,
			CreatedAt:    field.CreatedAt,
			UpdatedAt:    field.UpdatedAt,
		})
	}

	pagination := &util.PaginationParam{
		Count: total,
		Page:  param.Page,
		Limit: param.Limit,
		Data:  fieldResults,
	}

	response := util.GeneratePagination(*pagination)
	return &response, nil
}

func (fs *FieldService) GetAllWithoutPagination(ctx context.Context) ([]dto.FieldResponse, error) {
	fields, err := fs.repository.GetField().FindAllWithoutPagination(ctx)
	if err != nil {
		return nil, err
	}

	fieldResults := make([]dto.FieldResponse, 0, len(fields))
	for _, field := range fields {
		fieldResults = append(fieldResults, dto.FieldResponse{
			UUID:         field.UUID,
			Name:         field.Name,
			PricePerHour: field.PricePerHour,
			Images:       field.Images,
		})
	}

	return fieldResults, nil
}

func (fs *FieldService) GetByUUID(ctx context.Context, uuid string) (*dto.FieldResponse, error) {
	field, err := fs.repository.GetField().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	fieldResult := dto.FieldResponse{
		UUID:         field.UUID,
		Code:         field.Code,
		Name:         field.Name,
		PricePerHour: field.PricePerHour,
		Images:       field.Images,
		CreatedAt:    field.CreatedAt,
		UpdatedAt:    field.UpdatedAt,
	}

	return &fieldResult, nil
}

func (fs *FieldService) Create(ctx context.Context, req *dto.FieldRequest) (*dto.FieldResponse, error) {
	imageUrl, err := fs.uploadImage(ctx, req.Images)
	if err != nil {
		return nil, err
	}

	field, err := fs.repository.GetField().Create(ctx, &models.Field{
		Code:         req.Code,
		Name:         req.Name,
		PricePerHour: req.PricePerHour,
		Images:       imageUrl,
	})
	if err != nil {
		return nil, err
	}

	response := &dto.FieldResponse{
		UUID:         field.UUID,
		Code:         field.Code,
		Name:         field.Name,
		PricePerHour: field.PricePerHour,
		Images:       field.Images,
		CreatedAt:    field.CreatedAt,
		UpdatedAt:    field.UpdatedAt,
	}

	return response, nil
}

func (fs *FieldService) Update(ctx context.Context, uuid string, req *dto.UpdateFieldRequest) (*dto.FieldResponse, error) {
	field, err := fs.repository.GetField().FindByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	var imageUrl []string
	if req.Images == nil {
		imageUrl = field.Images
	} else {
		imageUrl, err = fs.uploadImage(ctx, req.Images)
		if err != nil {
			return nil, err
		}
	}

	fieldResult, err := fs.repository.GetField().Update(ctx, uuid, &models.Field{
		Code:         req.Code,
		Name:         req.Name,
		PricePerHour: req.PricePerHour,
		Images:       imageUrl,
	})
	if err != nil {
		return nil, err
	}

	return &dto.FieldResponse{
		UUID:         fieldResult.UUID,
		Code:         fieldResult.Code,
		Name:         fieldResult.Name,
		PricePerHour: fieldResult.PricePerHour,
		Images:       fieldResult.Images,
		CreatedAt:    fieldResult.CreatedAt,
		UpdatedAt:    fieldResult.UpdatedAt,
	}, nil
}

func (fs *FieldService) Delete(ctx context.Context, uuid string) error {
	_, err := fs.repository.GetField().FindByUUID(ctx, uuid)
	if err != nil {
		return err
	}

	err = fs.repository.GetField().Delete(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}

func (fs *FieldService) validateUpdload(images []multipart.FileHeader) error {
	if images == nil || len(images) == 0 {
		return errConstant.ErrInvalidUploadFile
	}

	for _, image := range images {
		if image.Size > 5*1024*1024 {
			return errConstant.ErrSizeTooBig
		}
	}

	return nil
}

func (fs *FieldService) processAndUploadImage(ctx context.Context, image multipart.FileHeader) (string, error) {
	file, err := image.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, file)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("images/%s-%s-%s", time.Now().Format("20060102150405"), image.Filename, path.Ext(image.Filename))
	url, err := fs.gcs.UploadFile(ctx, filename, buffer.Bytes())
	if err != nil {
		return "", err
	}

	return url, nil
}

func (fs *FieldService) uploadImage(ctx context.Context, images []multipart.FileHeader) ([]string, error) {
	err := fs.validateUpdload(images)
	if err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(images))
	for _, image := range images {
		url, err := fs.processAndUploadImage(ctx, image)
		if err != nil {
			return nil, err
		}

		urls = append(urls, url)
	}

	return urls, nil
}
