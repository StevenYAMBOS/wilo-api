package main

import (
	"github.com/StevenYAMBOS/wilo-api/config"
	"github.com/StevenYAMBOS/wilo-api/database"
	"github.com/StevenYAMBOS/wilo-api/internal/auth"
	"github.com/gin-gonic/gin"
)

func main() {
    config.LoadConfig()
    database.Connect()

    r := gin.Default()

    r.POST("/auth/register", auth.Register)
    r.POST("/auth/login", auth.Login)

    protected := r.Group("/")
    protected.Use(auth.AuthMiddleware())
    protected.GET("/dashboard", auth.Dashboard)

    r.Run(config.Port)
}
