package api

import (
    "github.com/gin-gonic/gin"
    "github.com/marceloneiva/myapi/internal/api/handlers"
)

func SetupRoutes() *gin.Engine {
    r := gin.Default()

    user := r.Group("/users")
    {
        user.GET("/", handlers.GetUsers)
    }

    return r
}