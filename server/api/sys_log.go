package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "nginx-ui/pkg/helper"
    "nginx-ui/pkg/logger"
    "nginx-ui/server/model"
)

func GetSysLog(c *gin.Context) {
    var req model.SysLogListReq
    if err := c.ShouldBindQuery(&req); err != nil {
        helper.FailWithMessage("query bind failed:"+err.Error(), c)
        return
    }

    data, err := model.GetSysLogs(c, req)
    if err != nil {
        logger.Errorf("failed to get sys_log: %v", err)
        helper.FailWithMessage("get sys_log failed:"+err.Error(), c)
        return
    }

    c.JSON(http.StatusOK, data)
    return
}
