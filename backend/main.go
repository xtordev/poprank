package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/",func (c *gin.Context){
		body,_:= io.ReadAll(c.Request.Body)
		println(string(body))
		c.JSON(http.StatusOK,gin.H{
			"message":"Welcome to Pop Rank API",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}