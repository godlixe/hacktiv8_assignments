package controller

import (
	"net/http"
	"order_service/common"
	"order_service/dto"
	"order_service/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController interface {
	GetOrders(ctx *gin.Context)
	GetOrderByID(ctx *gin.Context)
	CreateOrder(ctx *gin.Context)
	UpdateOrder(ctx *gin.Context)
	DeleteOrder(ctx *gin.Context)
}

type orderController struct {
	orderService service.OrderService
}

func NewOrderController(ors service.OrderService) OrderController {
	return &orderController{
		orderService: ors,
	}
}

func (c *orderController) GetOrders(ctx *gin.Context) {
	result, err := c.orderService.GetOrders(ctx)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get orders", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) GetOrderByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get order", "No id found", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.orderService.GetOrderByID(ctx, id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get order", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)

}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var orderDTO dto.OrderCreateDTO

	if err := ctx.ShouldBind(&orderDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind order", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.orderService.CreateOrder(ctx.Request.Context(), orderDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to create order", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) UpdateOrder(ctx *gin.Context) {
	var orderDTO dto.OrderUpdateDTO

	if err := ctx.ShouldBind(&orderDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind order", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.orderService.UpdateOrder(ctx, orderDTO)
	if err != nil {
		res := common.BuildErrorResponse("Failed to update order", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, response)
}

func (c *orderController) DeleteOrder(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		res := common.BuildErrorResponse("Failed to get id", "No id found", common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if err := c.orderService.DeleteOrder(ctx.Request.Context(), id); err != nil {
		res := common.BuildErrorResponse("Failed to delete order", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	response := common.BuildResponse(true, "OK", common.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
