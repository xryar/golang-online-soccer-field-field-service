package controllers

import (
	"field-service/services"

	"github.com/gin-gonic/gin"
)

type FieldController struct {
	service services.IRegistryService
}

type IFieldController interface {
	GetAllWithPagination(*gin.Context)
	GetAllWithoutPagination(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func NewFieldController(service services.IRegistryService) IFieldController {
	return &FieldController{service: service}
}

func (fc *FieldController) GetAllWithPagination(c *gin.Context) {}

func (fc *FieldController) GetAllWithoutPagination(c *gin.Context) {}

func (fc *FieldController) GetByUUID(c *gin.Context) {}

func (fc *FieldController) Create(c *gin.Context) {}

func (fc *FieldController) Update(c *gin.Context) {}

func (fc *FieldController) Delete(c *gin.Context) {}
