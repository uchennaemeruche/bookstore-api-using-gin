package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/uchennaemeruche/go-api-examples/ginApi/mvc/controllers"
	"github.com/uchennaemeruche/go-api-examples/ginApi/mvc/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading enviroment variables")
	}

	var config = models.DBConfig{
		DBName:     os.Getenv("DB_NAME"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbUser:     os.Getenv("DB_USER"),
		DbPort:     os.Getenv("DB_PORT"),
	}

	router := gin.Default()

	models.ConnectDatabase(models.DBConnectionString{
		Dialect: "mysql",
		Dsn:     fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DBName),
	})

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	// File upload
	router.MaxMultipartMemory = 8 << 20 //8 MiB
	// router.Static("/", "./public")
	router.POST("/uploads", controllers.UploadFile)

	router.Run()

}
