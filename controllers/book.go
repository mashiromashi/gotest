package controllers

import (
	"mashiromashi/gotest/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func FindBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"books": books})
}

func CreateBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validate input
	var input models.CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var book models.Book
	c.BindJSON(&book)

	db.Create(&book)
	c.JSON(http.StatusOK, gin.H{"book": book})
}
