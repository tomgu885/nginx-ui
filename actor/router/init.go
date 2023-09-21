package router

import (
	"github.com/gin-gonic/gin"
	"nginx-ui/actor/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.GET("/hello", func(c *gin.Context) {
		c.String(200, "world")
	})

	r.GET("/status", api.NodeStatus)
	return r
}
