package main

import (
	"order_service/config"
	"order_service/controller"
	"order_service/repository"
	"order_service/service"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	var (
		db              *gorm.DB                   = config.SetupDatabaseConnection()
		OrderRepository repository.OrderRepository = repository.NewOrderRepository(db)
		ItemRepository  repository.ItemRepository  = repository.NewItemRepository(db)
		orderService    service.OrderService       = service.NewOrderService(OrderRepository, ItemRepository)
		orderController controller.OrderController = controller.NewOrderController(orderService)
	)
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	orderRoutes := server.Group("/orders")
	{
		orderRoutes.GET("", orderController.GetOrders)
		orderRoutes.POST("/create", orderController.CreateOrder)
		orderRoutes.PUT("/update", orderController.UpdateOrder)
		orderRoutes.DELETE("/:id", orderController.DeleteOrder)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}
