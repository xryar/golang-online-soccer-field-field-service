package controllers

import (
	errValidation "field-service/common/error"
	"field-service/common/response"
	"field-service/domain/dto"
	"field-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FieldScheduleController struct {
	service services.IRegistryService
}

type IFieldScheduleController interface {
	GetAllWithPagination(*gin.Context)
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

func (fsc *FieldScheduleController) GetAllWithPagination(c *gin.Context) {
	var params dto.FieldScheduleRequestParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusBadRequest,
			Err:     err,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     c,
		})
		return
	}

	result, err := fsc.service.GetFieldSchedule().GetAllWithPagination(c, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Data: result,
		Gin:  c,
	})
}

func (fsc *FieldScheduleController) GetAllByFieldIDAndDate(c *gin.Context) {
	var params dto.FieldScheduleByFieldIDAndDateRequestParam
	err := c.ShouldBindQuery(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusBadRequest,
			Err:     err,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     c,
		})
		return
	}

	result, err := fsc.service.GetFieldSchedule().GetAllByFieldIDAndDate(c, c.Param("uuid"), params.Date)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Data: result,
		Gin:  c,
	})
}

func (fsc *FieldScheduleController) GetByUUID(c *gin.Context) {
	result, err := fsc.service.GetFieldSchedule().GetByUUID(c, c.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Data: result,
		Gin:  c,
	})
}

func (fsc *FieldScheduleController) Create(c *gin.Context) {
	var params dto.FieldScheduleRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusBadRequest,
			Err:     err,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     c,
		})
		return
	}

	err = fsc.service.GetFieldSchedule().Create(c, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusCreated,
		Gin:  c,
	})
}

func (fsc *FieldScheduleController) Update(c *gin.Context) {
	var params dto.UpdateFieldScheduleRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusBadRequest,
			Err:     err,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     c,
		})
		return
	}

	result, err := fsc.service.GetFieldSchedule().Update(c, c.Param("uuid"), &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusCreated,
		Data: result,
		Gin:  c,
	})
}

func (fsc *FieldScheduleController) UpdateStatus(c *gin.Context) {
	var request dto.UpdateStatusFieldScheduleRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusBadRequest,
			Err:     err,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     c,
		})
		return
	}

	err = fsc.service.GetFieldSchedule().UpdateStatus(c, &request)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Gin:  c,
	})
}

func (fsc *FieldScheduleController) Delete(c *gin.Context) {
	err := fsc.service.GetFieldSchedule().Delete(c, c.Param("uuid"))
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusOK,
		Gin:  c,
	})
}

func (fsc *FieldScheduleController) GenerateScheduleForOneMonth(c *gin.Context) {
	var params dto.GenerateFieldScheduleForOneMonthRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(params)
	if err != nil {
		errMessage := http.StatusText(http.StatusUnprocessableEntity)
		errResponse := errValidation.ErrValidationResponse(err)
		response.HttpResponse(response.ParamHTTPResponse{
			Code:    http.StatusBadRequest,
			Err:     err,
			Message: &errMessage,
			Data:    errResponse,
			Gin:     c,
		})
		return
	}

	err = fsc.service.GetFieldSchedule().GenerateScheduleForOneMonth(c, &params)
	if err != nil {
		response.HttpResponse(response.ParamHTTPResponse{
			Code: http.StatusBadRequest,
			Err:  err,
			Gin:  c,
		})
		return
	}

	response.HttpResponse(response.ParamHTTPResponse{
		Code: http.StatusCreated,
		Gin:  c,
	})
}
