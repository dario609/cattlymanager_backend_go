package activitiesapi

import (
	"net/http"
	"strconv"

	"livestock-saas/server/internal/activities"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

func GetActivity(c *gin.Context) {
	id := c.Param("id")

	var activity activities.Activity
	if err := database.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func ListActivities(c *gin.Context) {
	// Optional query parameter for filtering by type (e.g., ?type=dead)
	activityType := c.Query("type")

	var activities []activities.Activity
	query := database.DB.Preload("Animal").Order("date DESC")

	if activityType != "" {
		query = query.Where("type = ?", activityType)
	}

	if err := query.Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch activities"})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func ListActivitiesByAnimal(c *gin.Context) {
	animalIDStr := c.Param("animal_id")
	animalID, err := strconv.ParseUint(animalIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid animal_id"})
		return
	}

	var activities []activities.Activity
	if err := database.DB.Where("animal_id = ?", uint(animalID)).Order("date DESC").Find(&activities).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch activities"})
		return
	}

	c.JSON(http.StatusOK, activities)
}
