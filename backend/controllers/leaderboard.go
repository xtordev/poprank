package controllers

import (
	"backend/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func allLeaderboard(c *gin.Context) {
	var ranking []db.Student
	db.GormDB.Order("Elo DESC").Limit(10).Find(&ranking)
	c.JSON(http.StatusOK, ranking)
}

func standingLeaderboard(c *gin.Context) {
	var APIObj struct {
		Code string `binding:"required"`
	}
	if err := c.ShouldBind(&APIObj); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var student db.Student
	if res := db.GormDB.Where("code = ?", APIObj.Code).First(&student); res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	var ranking int64
	err := db.GormDB.Raw(`
        SELECT COUNT(*)
        FROM students
        WHERE Elo > ?
        OR (Elo = ? AND Code < ?)
    `, student.Elo, student.Elo, student.Code).Scan(&ranking).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate ranking"})
		return
	}
	ranking++

	var totalStudents, studentsBelow int64
	err = db.GormDB.Raw(`
        SELECT COUNT(*)
        FROM students
    `).Scan(&totalStudents).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate total students"})
		return
	}

	err = db.GormDB.Raw(`
        SELECT COUNT(*)
        FROM students
        WHERE Elo < ?
    `, student.Elo).Scan(&studentsBelow).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate students below"})
		return
	}

	percentile := 100.0 * float64(studentsBelow) / float64(totalStudents)

	c.JSON(http.StatusOK, gin.H{
		"student":    student,
		"ranking":    ranking,
		"percentile": percentile,
	})

}

func LeaderboardRoutes(r *gin.Engine) {
	routes := r.Group("/leaderboard")
	{
		routes.GET("", allLeaderboard)
		routes.POST("/standing", standingLeaderboard)
	}
}
