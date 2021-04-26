package main

import (
	"log"

	"github.com/gin-gonic/gin"
	db "github.com/rickseven/dockerized-golang-mysql-api-sample/db"
	product "github.com/rickseven/dockerized-golang-mysql-api-sample/services"
)

func main() {
	log.Println("Starting server..")

	db.Init()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.GET("/", product.GetAll)
			products.POST("/", product.Create)
		}
	}

	r.Run(":8081")
}
