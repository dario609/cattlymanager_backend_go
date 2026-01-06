package animalsapi

import (
	"livestock-saas/server/internal/animals"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

func GetAnimal(c *gin.Context) {
	id := c.Param("id")
	var animal animals.Animal

	if err := database.DB.First(&animal, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Animal not found"})
		return
	}

	c.JSON(200, animal)
}
