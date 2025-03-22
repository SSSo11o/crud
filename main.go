package main

import (
	database "CRUD/bd"
	"CRUD/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.ConnectDB()

	r := gin.Default()
	r.POST("/url", handler.CreateUrl)
	r.GET("/url/:id", handler.GetUrl)
	r.PUT("/url/:id", handler.UpdateUrl)
	r.DELETE("/url/:id", handler.DeleteUrl)

	log.Println("Сервер запущен на порту 8080...")
	r.Run(":8080")
}

//
