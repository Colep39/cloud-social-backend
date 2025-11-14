package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestSignup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "signup works!"})
}
