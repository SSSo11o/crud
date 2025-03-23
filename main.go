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
	r.GET("/url/:id", handler.GetUrlByID)
	r.PUT("/url/:id", handler.UpdateUrl)
	r.DELETE("/url/:id", handler.DeleteUrl)
	r.GET("/url", handler.GetAllUrls)

	log.Println("Сервер запущен на порту 8080...")
	r.Run(":8080")
}
