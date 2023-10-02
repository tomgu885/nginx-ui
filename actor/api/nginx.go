package api

import (
    "github.com/gin-gonic/gin"
    "nginx-ui/actor/model"
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

func Updates(c *gin.Context) {
    var req model.UpdateReq
    if err := c.ShouldBindJSON(req); err != nil {
        helper.FailWithMessage("json错误", c)
        return
    }

    requestId := c.GetString("x-request-id")

    go func(force bool, requestId string) {
        services.ServerConfigReload(force)

    }(req.Force == 1, requestId)

    helper.OkWithMessage("更新中", c)

}
