package api

import (
	"github.com/gin-gonic/gin"
	"nginx-ui/actor/services"
	"nginx-ui/pkg/helper"

	"strconv"
)

// @Router
func Update(c *gin.Context) {
	siteId := c.Query("id")

	id, err := strconv.Atoi(siteId)
	if err != nil {
		helper.FailWithMessage("id不是数字", c)
		return
	}

	conf, err := services.ServerConfig(uint(id))
	c.String(200, conf)
	return
}
