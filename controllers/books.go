package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/uchennaemeruche/go-api-examples/ginApi/mvc/models"
)

func FindBooks(c *gin.Context) {

	var books []models.Book

	// models.DB.Set("gorm:auto_preload", true).Find(&books)
	// models.DB.Preload("Authors").Find(&books)
	models.DB.Preload("Authors", func(db *gorm.DB) *gorm.DB {
		fmt.Println(db.Select("id", "Name"))
		return db.Select("id", "Name")
	}).Find(&books)
	// models.DB.Joins("Authors").Find(&books)
	// models.DB.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})

}

func FindBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func CreateBook(c *gin.Context) {
	// Validate input
	var input models.Book
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{Title: input.Title, Authors: input.Authors}

	models.DB.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func UpdateBook(c *gin.Context) {

	var book models.Book

	// check if book with id exist
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate input

	var input models.UpdateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	// check if book with id exist
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})

}

func DeleteAllBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Delete(&books)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
