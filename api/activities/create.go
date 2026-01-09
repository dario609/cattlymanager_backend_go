package activitiesapi

import (
	"net/http"
	"time"

	"livestock-saas/server/internal/activities"
	"livestock-saas/server/internal/animals"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type createActivityRequest struct {
	AnimalID     uint    `json:"animal_id" binding:"required"`
	Type         string  `json:"type" binding:"required,oneof=sale dead"`
	Date         string  `json:"date" binding:"required"`
	SalePrice    float64 `json:"sale_price,omitempty"`
	CustomerName string  `json:"customer_name,omitempty"`
	DeathCause   string  `json:"death_cause,omitempty"`
	Notes        string  `json:"notes,omitempty"`
}

func CreateActivity(c *gin.Context) {
	var req createActivityRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate required fields based on activity type
	errs := make(map[string]string)

	if req.Type == "sale" && req.SalePrice == 0 {
		errs["sale_price"] = "sale_price is required for sale activities"
	}

	if req.Type == "dead" && req.DeathCause == "" {
		errs["death_cause"] = "death_cause is required for dead activities"
	}

	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	// Parse date
	activityDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
		return
	}

	activity := &activities.Activity{
		AnimalID:     req.AnimalID,
		Type:         req.Type,
		Date:         activityDate,
		SalePrice:    req.SalePrice,
		CustomerName: req.CustomerName,
		DeathCause:   req.DeathCause,
		Notes:        req.Notes,
	}

	// Create activity and update animal status in a transaction
	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(activity).Error; err != nil {
			return err
		}

		var newStatus string
		if req.Type == "sale" {
			newStatus = "Sold"
		} else if req.Type == "dead" {
			newStatus = "Dead"
		}

		if newStatus != "" {
			if err := tx.Model(&animals.Animal{}).Where("id = ?", req.AnimalID).Update("status", newStatus).Error; err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create activity or update animal status"})
		return
	}

	c.JSON(http.StatusCreated, activity)
}
