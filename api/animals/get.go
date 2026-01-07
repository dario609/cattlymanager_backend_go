package animalsapi

import (
	"livestock-saas/server/internal/animals"
	"livestock-saas/server/internal/database"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAnimal(c *gin.Context) {
	id := c.Param("id")
	var animal animals.Animal

	if err := database.DB.First(&animal, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Animal not found"})
		return
	}

	log.Printf("GetAnimal: returning id=%v ear_tag_color=%q", animal.ID, animal.EarTagColor)

	c.JSON(200, animal)
}
