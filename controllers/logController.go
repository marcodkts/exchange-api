package controllers

import (
	"exchange-api/initializers"
	"exchange-api/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewLog(userID uint, method string, path string, status int) *models.Log {
	return &models.Log{
		UserID:    userID,
		Timestamp: time.Now(),
		Method:    method,
		Path:     path,
		Status:    status,
	}
}

func LogRequest(userId uint, method string, path string, status int) {
    log := NewLog(userId, method, path, status)
    initializers.DB.Create(log)
}


func LogIndex(c *gin.Context) {
	var logs []models.Log
	initializers.DB.Find(&logs)

	c.JSON(http.StatusOK, gin.H{
		"logs": logs,
	})
}

func LogShow(c *gin.Context) {
	id := c.Param("id")

	var log models.Log
	initializers.DB.Find(&log, id)

	if log.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Log not found.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"log": log,
	})
}