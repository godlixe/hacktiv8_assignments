package main

import (
	"order_service/config"
	"order_service/controller"
	"order_service/repository"
	"order_service/routes"
	"order_service/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	var (
		db               *gorm.DB                    = config.SetupDatabaseConnection()
		OrderRepository  repository.OrderRepository  = repository.NewOrderRepository(db)
		ItemRepository   repository.ItemRepository   = repository.NewItemRepository(db)
		PersonRepository repository.PersonRepository = repository.NewPersonRepository(os.Getenv("GIDHAN_API"))

		orderService service.OrderService = service.NewOrderService(OrderRepository, ItemRepository, PersonRepository)

		orderController controller.OrderController = controller.NewOrderController(orderService)
	)
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	routes.OrderRoutes(server, orderController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)

}
