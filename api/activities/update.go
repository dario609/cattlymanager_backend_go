package activitiesapi

import (
	"net/http"
	"time"

	"livestock-saas/server/internal/activities"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

type updateActivityRequest struct {
	Type         string  `json:"type,omitempty"`
	Date         string  `json:"date,omitempty"`
	SalePrice    float64 `json:"sale_price,omitempty"`
	CustomerName string  `json:"customer_name,omitempty"`
	DeathCause   string  `json:"death_cause,omitempty"`
	Notes        string  `json:"notes,omitempty"`
}

func UpdateActivity(c *gin.Context) {
	id := c.Param("id")

	var activity activities.Activity
	if err := database.DB.First(&activity, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}

	var req updateActivityRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if provided
	if req.Type != "" {
		activity.Type = req.Type
	}
	if req.Date != "" {
		activityDate, err := time.Parse("2006-01-02", req.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
			return
		}
		activity.Date = activityDate
	}
	if req.SalePrice != 0 {
		activity.SalePrice = req.SalePrice
	}
	if req.CustomerName != "" {
		activity.CustomerName = req.CustomerName
	}
	if req.DeathCause != "" {
		activity.DeathCause = req.DeathCause
	}
	if req.Notes != "" {
		activity.Notes = req.Notes
	}

	if err := database.DB.Save(&activity).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update activity"})
		return
	}

	c.JSON(http.StatusOK, activity)
}
