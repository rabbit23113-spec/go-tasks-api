package handler

import (
	"main/package/service"
	"main/package/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	auth service.AuthService
}

func (ah *AuthHandler) SignUp(c *gin.Context) {
	var dto types.SignUpDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if err := ah.auth.SignUp(dto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"status": "signed up successfully"})
}

func (ah *AuthHandler) SignIn(c *gin.Context) {
	var dto types.SignInDto
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	if err := ah.auth.SignIn(dto); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"status": "signed in successfully"})
}
