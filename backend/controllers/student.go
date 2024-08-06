package controllers

import (
	"backend/db"
	"backend/middlewares"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func isProduction(c *gin.Context) {
	if os.Getenv("PRODUCTION") == "true" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		c.Abort()
		return
	}
}
func findAll(c *gin.Context) {

	var students []db.Student
	db.GormDB.Find(&students)
	c.JSON(http.StatusOK, students)
}
func findOne(c *gin.Context) {
	var student db.Student
	id := c.Param("id")
	result := db.GormDB.First(&student, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}
func create(c *gin.Context) {
	var student db.Student
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	codeExist := db.GormDB.Where("code = ?", student.Code).First(&student)
	if codeExist.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Student Code Already Exist"})
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file provided"})
		return
	}
	student.Image, err = middlewares.UploadImage(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error uploading image: " + err.Error()})
		return
	}
	if err := db.GormDB.Create(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, student)
}
func update(c *gin.Context) {
	var student db.Student
	id := c.Param("id")
	result := db.GormDB.First(&student, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	if err := c.ShouldBind(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	codeExist := db.GormDB.Where("code = ?", student.Code).First(&student)
	if codeExist.Error == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Student Code Already Exist"})
		return
	}
	file, _ := c.FormFile("file")
	if file != nil {
		middlewares.DeleteImage(student.Image)
		student.Image, _ = middlewares.UploadImage(file)
	}
	if err := db.GormDB.Save(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"student": student})
}
func delete(c *gin.Context) {
	var student db.Student
	id := c.Param("id")
	result := db.GormDB.First(&student, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	middlewares.DeleteImage(student.Image)
	if err := db.GormDB.Delete(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, student)
}

func StudentsRoutes(r *gin.Engine) {
	routes := r.Group("/students")
	{
		routes.GET("/", isProduction, findAll)
		routes.GET("/:id", isProduction, findOne)
		routes.POST("/", isProduction, create)
		routes.PUT("/:id", isProduction, update)
		routes.DELETE("/:id", isProduction, delete)
	}
}
