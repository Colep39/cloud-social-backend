package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/colep39/cloud-social-backend/services/auth/internal/auth"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func Signup(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SignupRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Hash password
		hash, err := auth.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}

		// Insert to DB
		_, err = db.Exec(
			context.Background(),
			`INSERT INTO users (email, password_hash) VALUES ($1, $2)`,
			req.Email,
			hash,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}
