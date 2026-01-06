package usersapi

import (
	"livestock-saas/server/internal/database"
	"livestock-saas/server/internal/users"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRoutes(r *gin.Engine) {
	repo := users.NewRepository(database.DB)

	r.POST("/users", func(c *gin.Context) {
		var input users.User
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to secure password"})
			return
		}

		input.Password = string(hashed)

		if err := repo.Create(&input); err != nil {
			c.JSON(http.StatusConflict, gin.H{"error": "user already exists"})
			return
		}

		c.JSON(http.StatusCreated, input)
	})

	r.GET("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")

		id, err := strconv.ParseUint(idParam, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": "invalid user id"})
			return
		}

		user, err := repo.FindByID(uint(id))
		if err != nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}

		c.JSON(200, user)
	})

}
