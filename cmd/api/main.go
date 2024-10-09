package main

import (
    "github.com/gin-gonic/gin"
    "github.com/buya-v/gsdf/config"
    "github.com/buya-v/gsdf/internal/db"
    "github.com/buya-v/gsdf/internal/auth"
    "github.com/buya-v/gsdf/internal/user"
)

func main() {
    config.LoadConfig()
    db.ConnectDatabase()

    r := gin.Default()
    auth.RegisterRoutes(r)
    user.RegisterRoutes(r)
    r.Run()
}
