package animalsapi

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	animals := r.Group("/animals")
	{
		animals.POST("/", CreateAnimal)
		animals.GET("/", ListAnimals)
		animals.GET("/:id", GetAnimal)
		animals.PUT("/:id", UpdateAnimal)
		animals.DELETE("/:id", DeleteAnimal)
	}
}
