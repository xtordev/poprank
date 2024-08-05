package router

import (
	"backend/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {
	routes := r.Group("/")
	{
		routes.GET("/", func(c *gin.Context) {
			// body, _ := io.ReadAll(c.Request.Body)
			// println(string(body))
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to Pop Rank API",
			})
		})
	}
	controllers.StudentsRoutes(r)
}
