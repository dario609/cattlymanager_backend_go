package auth

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"livestock-saas/server/internal/database"
	"livestock-saas/server/internal/users"
)

var jwtSecret = []byte("CHANGE_ME_LATER")

func RegisterRoutes(r *gin.Engine) {
	repo := users.NewRepository(database.DB)

	r.POST("/login", func(c *gin.Context) {
		var req LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := repo.FindByEmail(req.Email)
		if err != nil {
			log.Printf("auth: FindByEmail error for %s: %v", req.Email, err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid Email"})
			return
		}
		log.Printf("Email and Password: %s, %s", req.Email, req.Password)
		log.Printf("User Password: %s", user.Password)
		if err := bcrypt.CompareHashAndPassword(
			[]byte(user.Password),
			[]byte(req.Password),
		); err != nil {
			log.Printf("auth: password mismatch for %s: %v", req.Email, err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":  user.ID,
			"farm": user.FarmID,
			"role": user.Role,
			"exp":  time.Now().Add(24 * time.Hour).Unix(),
		})

		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	})
}
