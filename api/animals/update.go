package animalsapi

import (
	"livestock-saas/server/internal/animals"
	"livestock-saas/server/internal/database"

	"github.com/gin-gonic/gin"
)

func UpdateAnimal(c *gin.Context) {
	id := c.Param("id")
	var animal animals.Animal

	if err := database.DB.First(&animal, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Animal not found"})
		return
	}

	var body animals.Animal
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid data"})
		return
	}

	database.DB.Model(&animal).Updates(body)

	c.JSON(200, animal)
}
