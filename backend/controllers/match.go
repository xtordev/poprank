package controllers

import (
	"backend/db"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func genMatch(c *gin.Context) {
	var student1, student2 db.Student

	// Query to select a random student
	err := db.GormDB.Raw("SELECT * FROM students ORDER BY RANDOM() LIMIT 1").Scan(&student1).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Query to select a different random student
	err = db.GormDB.Raw("SELECT * FROM students WHERE id != $1 ORDER BY RANDOM() LIMIT 1", student1.ID).Scan(&student2).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"student1": student1,
		"student2": student2,
	})
}

const kFactor = 32.0 // Adjust this value as needed

func calculateElo(rating1, rating2, winner int) (float64, float64) {
	expectedScore1 := 1.0 / (1.0 + math.Pow(10, float64(rating2-rating1)/400.0))
	expectedScore2 := 1.0 - expectedScore1

	newRating1 := float64(rating1) + kFactor*(float64(winner)-expectedScore1)
	newRating2 := float64(rating2) + kFactor*(float64(1-winner)-expectedScore2)

	return newRating1, newRating2
}

type API struct {
	StudentId1  int `binding:"required"`
	StudentId2  int `binding:"required"`
	MatchWinner int `binding:"required"`
}

func matchWin(c *gin.Context) {
	var student1 db.Student
	var student2 db.Student
	var bodyJson API
	if err := c.ShouldBind(&bodyJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if res := db.GormDB.First(&student1, bodyJson.StudentId1); res.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Student1 Id doesn't exist"})
		return
	}

	if res := db.GormDB.First(&student2, bodyJson.StudentId2); res.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Student2 Id doesn't exist"})
		return
	}
	if bodyJson.MatchWinner == bodyJson.StudentId1 {
		newRating1, newRating2 := calculateElo(student1.Elo, student2.Elo, 1)
		student1.Elo = int(newRating1)
		student2.Elo = int(newRating2)

		db.GormDB.Save(&student1)
		db.GormDB.Save(&student2)
	} else if bodyJson.MatchWinner == bodyJson.StudentId2 {
		newRating2, newRating1 := calculateElo(student2.Elo, student1.Elo, 1)
		student1.Elo = int(newRating1)
		student2.Elo = int(newRating2)

		db.GormDB.Save(&student1)
		db.GormDB.Save(&student2)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MatchWinner must be either StudentId1 or StudentId2"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "Successfully updated Elo",
		"student1": student1,
		"student2": student2,
	})
}

func MatchRoutes(r *gin.Engine) {
	routes := r.Group("/matches")
	{
		routes.GET("", genMatch)
		routes.POST("/win", matchWin)
	}
}
