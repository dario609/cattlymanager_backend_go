package animalsapi

import (
	"livestock-saas/server/internal/animals"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

func DeleteAnimal(c *gin.Context) {
	id := c.Param("id")
	var animal animals.Animal

	if err := database.DB.First(&animal, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Animal not found"})
		return
	}

	if err := database.DB.Delete(&animal).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(200, gin.H{"message": "Animal removed"})
}
