package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "nginx-ui/pkg/helper"
    "nginx-ui/server/model"
)

// cert & sites
func Updates(c *gin.Context) {
    // @todo sign, valid
    var req model.SiteListReq
    if err := c.ShouldBindQuery(&req); err != nil {
        helper.FailWithMessage(fmt.Sprintf("参数绑定失败:%v", err), c)
        return
    }
    fmt.Println("keyword:", req.Keyword)
    fmt.Println("ids:", req.Ids)
    req.PageSize = 1999
    req.WithCert = true
    sites, total, err := model.GetSites(req)

    if err != nil {
        helper.FailWithMessage("获取失败", c)
        return
    }

    helper.OkWithDetailed(helper.PageResult{
        List:     sites,
        Total:    total,
        Page:     req.Page,
        PageSize: req.PageSize,
    }, "获取成功", c)
    return
}
