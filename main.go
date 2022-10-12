package main

import (
	"example/Tugas/config"
	"example/Tugas/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.DBInit()
	repoOrder := handler.NewRepoOrder(db)
	route := gin.Default()
	route.GET("/order", repoOrder.GetOrder)
	route.GET("/order/:id", repoOrder.FindById)

	route.POST("/order", repoOrder.CreateOrder)
	route.PUT("/order/:id", repoOrder.CreateOrder)
	route.DELETE("/order/:id", repoOrder.CreateOrder)
	route.Run(":8080")
}
