package handler

import (
	"CRUD/bd"
	"CRUD/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
)

func CreateUrl(c *gin.Context) {
	var url models.Url
	log.Println("Начало обработки запроса для создания")

	if err := c.ShouldBindJSON(&url); err != nil {
		log.Println("Ошибка при привязке JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось создать"})
		return
	}

	log.Printf("Получение данных для создания URL: %+v", url)
	result := database.DB.Create(&url)
	if result.Error != nil {
		log.Printf("Ошибка при создании URL: %v", result.Error)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не удалось создать"})
		return
	}

	log.Printf("Успешно создан URL: %+v", url)
	c.JSON(http.StatusOK, url)
}

func GetUrl(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("неверные параметр id: %v", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверные параметр"})
		return
	}
	log.Printf("Поиск url с id: %v", id)
	var url models.Url
	result := database.DB.First(&url, "id =? and is_active =?", id, true)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("Url с ID %d не найдён", id)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Url с id не найдён"})
		} else {
			log.Printf("Ошибка при извличение url: %v", result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось извлечь url"})
		}
		return
	}
	log.Printf("Найдено url с id: %v", id)
	c.JSON(http.StatusOK, url)
}

func UpdateUrl(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("Ошибка при приобразования id: %v", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не верные формат id"})
		return
	}
	id64 := int64(id)
	var url models.Url
	result := database.DB.First(&url, id64)
	if err := c.ShouldBindJSON(&url); err != nil {
		if result.Error != nil {
			log.Printf("Запись с id %d не найдён", id64)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Запись не найден"})
		} else {
			log.Printf("Ошибка при поиске записи с id %d: %v", id64, result.Error)
		}
		return
	}

	url.ID = id64
	result = database.DB.Model(&url).Updates(url)
	if result.Error != nil {
		log.Printf("Ошибка при обновление записы с ID %d: %v", id64, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось обновить запись"})
		return
	}

	log.Printf("Запись с Id %d успешно обновлена", id64)
	c.JSON(http.StatusOK, url)
}

func DeleteUrl(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("не верные параметр id: %v", id)
		c.JSON(http.StatusBadRequest, gin.H{"error": "не верные параметр"})
		return
	}
	log.Printf("Попытка удалить url c id %d", id)
	result := database.DB.Model(&models.Url{}).Where("id=?", id).Update("is_active", false).Update("updated_at", time.Now())
	if result.Error != nil {
		log.Printf("Ошибка при удалении с url c id %d: %v", id, result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении "})
		return
	}

	if result.RowsAffected == 0 {
		log.Printf(" url c id %d не найден", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "url не найден"})
		return
	}

	log.Printf("Успешно удалено url c id %d", id)
	c.JSON(http.StatusOK, nil)
}
