package main

import (
	"example/main.go/config"
	"example/main.go/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	repoOrder := handler.NewRepoOrder(db)
	route := gin.Default()
	route.GET("/order", repoOrder.GetOrder)
	route.GET("/order/:id", repoOrder.FindById)

	route.POST("/order", repoOrder.CreateOrder)
	route.PUT("/order/:id", repoOrder.UpdateOrder)
	route.DELETE("/order/:id", repoOrder.DeleteOrder)
	route.Run(":8080")
}
