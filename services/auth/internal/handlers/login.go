package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/colep39/cloud-social-backend/services/auth/internal/auth"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func Login(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login credentials"})
			return
		}

		// Fetch user
		var userID, storedHash string
		err := db.QueryRow(
			context.Background(),
			`SELECT id, password_hash FROM users WHERE email=$1`,
			req.Email,
		).Scan(&userID, &storedHash)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Compare password
		if !auth.CheckPassword(storedHash, req.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}

		// Generate JWT
		token, err := auth.GenerateToken(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
