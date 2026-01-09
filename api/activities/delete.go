package activitiesapi

import (
	"net/http"

	"livestock-saas/server/internal/activities"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

func DeleteActivity(c *gin.Context) {
	id := c.Param("id")

	var activity activities.Activity
	if err := database.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}

	if err := database.DB.Delete(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete activity"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "activity deleted successfully"})
}
