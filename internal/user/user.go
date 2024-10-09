package user

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func RegisterRoutes(r *gin.Engine) {
    r.GET("/users", getUsers)
}

func getUsers(c *gin.Context) {
    // Fetch users logic
    c.JSON(http.StatusOK, gin.H{"message": "Fetched users successfully"})
}
