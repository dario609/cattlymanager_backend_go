package animalsapi

import (
	"livestock-saas/server/internal/animals"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

func ListAnimals(c *gin.Context) {
	var list []animals.Animal

	if err := database.DB.Order("id DESC").Find(&list).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to load animals"})
		return
	}

	c.JSON(200, list)
}
