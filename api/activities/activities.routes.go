package activitiesapi

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	activities := r.Group("/activities")
	{
		activities.POST("", CreateActivity)
		activities.GET("", ListActivities)              // List all activities with optional type filter
		activities.GET("/animal/:animal_id", ListActivitiesByAnimal)
		activities.GET("/:id", GetActivity)            // Must be after specific routes
		activities.PUT("/:id", UpdateActivity)
		activities.DELETE("/:id", DeleteActivity)
	}
}
