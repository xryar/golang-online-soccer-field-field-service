package services

import (
	"field-service/common/response"
	"field-service/domain/dto"
	"field-service/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimeController struct {
	service services.IRegistryService
}

type ITimeController interface {
	GetAll(*gin.Context)
	GetByUUID(*gin.Context)
	Create(*gin.Context)
}

func NewTimeController(service services.IRegistryService) ITimeController {
	return &TimeController{service: service}
}

func (tc *TimeController) GetAll(c *gin.Context) {
	result, err := tc.service.GetTime().GetAll(c)
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

func (tc *TimeController) GetByUUID(c *gin.Context) {
	result, err := tc.service.GetTime().GetByUUID(c, c.Param("uuid"))
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

func (tc *TimeController) Create(c *gin.Context) {
	var request dto.TimeRequest
	result, err := tc.service.GetTime().Create(c, &request)
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
