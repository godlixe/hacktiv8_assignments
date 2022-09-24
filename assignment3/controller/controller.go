package controller

import (
	"assignment3/helpers"
	"assignment3/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetData(ctx *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(s service.Service) Controller {
	return &controller{
		service: s,
	}
}

func (c *controller) GetData(ctx *gin.Context) {
	result, err := c.service.GetData()
	if err != nil {
		response := helpers.BuildErrorResponse("Failed to get order", "No id found", helpers.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)
}
