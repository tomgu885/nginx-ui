package api

import "github.com/gin-gonic/gin"

type StatusApi struct {
}

func (StatusApi) Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
		"message": "hello from actor",
	})
}
