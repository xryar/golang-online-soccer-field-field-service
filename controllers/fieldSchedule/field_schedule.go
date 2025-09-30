package controllers

import (
	"field-service/services"

	"github.com/gin-gonic/gin"
)

type FieldScheduleController struct {
	service services.IRegistryService
}

type IFieldScheduleController interface {
	GetAllPagination(*gin.Context)
	GetAllByFieldIDAndDate(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)
	Update(*gin.Context)
	UpdateStatus(*gin.Context)
	Delete(*gin.Context)
	GenerateScheduleForOneMonth(*gin.Context)
}

func NewFieldScheduleController(service services.IRegistryService) IFieldScheduleController {
	return &FieldScheduleController{service: service}
}

func (fsc *FieldScheduleController) GetAllPagination(c *gin.Context) {}

func (fsc *FieldScheduleController) GetAllByFieldIDAndDate(c *gin.Context) {}

func (fsc *FieldScheduleController) GetByUUID(c *gin.Context) {}

func (fsc *FieldScheduleController) Create(c *gin.Context) {}

func (fsc *FieldScheduleController) Update(c *gin.Context) {}

func (fsc *FieldScheduleController) UpdateStatus(c *gin.Context) {}

func (fsc *FieldScheduleController) Delete(c *gin.Context) {}

func (fsc *FieldScheduleController) GenerateScheduleForOneMonth(c *gin.Context) {}
