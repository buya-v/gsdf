package auth

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterRoutes(r *gin.Engine) {
    r.POST("/register", register)
    r.POST("/login", login)
}

func register(c *gin.Context) {
    // Registration logic
    c.JSON(http.StatusOK, gin.H{"message": "User registered 
successfully"})
}

func login(c *gin.Context) {
    // Login logic
    c.JSON(http.StatusOK, gin.H{"message": "User logged in successfully"})
}
