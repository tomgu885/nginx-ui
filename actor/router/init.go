package router

import (
    "github.com/gin-gonic/gin"
    "nginx-ui/actor/api"
    "nginx-ui/pkg/middleware"
)

func InitRouter() *gin.Engine {
    r := gin.New()
    r.Use(gin.Logger())
    r.Use(middleware.XRequestID())
    r.GET("/", func(c *gin.Context) {
        c.String(200, "nginx-bl tiny cdn manager 3")
    })
    r.GET("/hello", func(c *gin.Context) {
        c.String(200, "world")
    })

    r.GET("/status", api.NodeStatus)

    r.POST("/updates", api.Updates)
    r.POST("")

    return r
}
