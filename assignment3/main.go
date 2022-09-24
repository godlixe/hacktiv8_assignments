package main

import (
	"assignment3/controller"
	"assignment3/repository"
	"assignment3/service"
	"os"
	"time"

	"github.com/gin-gonic/gin"
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
	var (
		repository repository.Repository = repository.NewDataRepository("database/db.json")
		service    service.Service       = service.NewService(repository)
		controller controller.Controller = controller.NewController(service)
	)
	server := gin.Default()

	go DoEvery(service.UpdateData, 15)

	server.GET("/", controller.GetData)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
