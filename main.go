package main

import (
	db "jbh-product/db"
	product "jbh-product/services"
	"log"

	"github.com/gin-gonic/gin"
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
