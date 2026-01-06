package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	animalsapi "livestock-saas/server/api/animals"
	authapi "livestock-saas/server/api/auth"
	usersapi "livestock-saas/server/api/users"

	"livestock-saas/server/internal/database"
)

func main() {
	godotenv.Load()

	database.Connect()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	usersapi.RegisterRoutes(r)
	authapi.RegisterRoutes(r)
	animalsapi.RegisterRoutes(r)

	r.Run(":8080")
}
