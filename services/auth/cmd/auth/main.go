package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/colep39/cloud-social-backend/services/auth/internal/db"
  "github.com/colep39/cloud-social-backend/services/auth/internal/handlers"

)

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/social?sslmode=disable")

	pool := db.Connect()
	defer pool.Close()

	r := gin.Default()

  r.POST("/signup", handlers.Signup)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "auth service healthy"})
	})

	log.Println("Auth service running on port 8081")
	if err := r.Run(":8081"); err != nil {
		log.Fatal(err)
	}
}

