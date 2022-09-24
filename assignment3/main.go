package main

import (
	"assignment3/controller"
	"assignment3/repository"
	"assignment3/service"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func DoEvery(function interface{}, duration time.Duration) {
	for {
		<-time.After(duration * time.Second)

		// type assert to run update function specifically
		update, ok := function.(func() error)
		if ok {
			go update()
		}
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Failed to load env,", err)
	}
	var (
		repository repository.Repository = repository.NewDataRepository(os.Getenv("FILEPATH"))
		service    service.Service       = service.NewService(repository)
		controller controller.Controller = controller.NewController(service)
	)
	server := gin.Default()

	// create db.json file
	os.OpenFile(os.Getenv("FILEPATH"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)

	// run goroutine
	go DoEvery(service.UpdateData, 1)

	// GET data endpoint
	server.GET("/", controller.GetData)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
