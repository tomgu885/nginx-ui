package api

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"
    "net/http"
    "nginx-ui/actor/services"
    "nginx-ui/pkg/helper"
    "nginx-ui/pkg/logger"
    "nginx-ui/server/model"
    "strconv"
)

// GetSites 获得站点列表
// @Summary 站点列表
// @Param domain query 域名
// @Router /api/sites [GET]
func GetSites(c *gin.Context) {
    var req model.SiteListReq
    if err := c.ShouldBindQuery(&req); err != nil {
        return
    }
    data, err := model.GetSites(c, req)
    if err != nil {
        logger.Errorf("GetSites faieled: %v", err)
        helper.FailWithMessage("获取数据失败", c)
        return
    }

    c.JSON(http.StatusOK, data)
}

// GetSite
// http://localhost:5173/api/sites/1
func GetSite(c *gin.Context) {
    id := cast.ToUint(c.Param("id"))
    site, err := model.GetSiteById(id)
    if err != nil {
        logger.Errorf("GetSites faieled: %v", err)
        helper.FailWithMessage("获取数据失败", c)
        return
    }

    c.JSON(http.StatusOK, site)
}

// EditSite 更新站点
func EditSite(c *gin.Context) {
    siteId := cast.ToUint(c.Param("id"))
    site, err := model.GetSiteById(siteId)
    if err != nil {
        helper.FailWithMessage("查询站点失败", c)
        return
    }

    if site.ID == 0 {
        helper.FailWithMessage("查询站点失败", c)
        return
    }

    var req model.SiteUpdateReq
    if errB := c.ShouldBindQuery(&req); errB != nil {
        helper.FailWithMessage("绑定json失败:"+errB.Error(), c)
        return
    }

    req.ID = siteId

    err = model.UpdateSite(req)
    if err != nil {
        logger.Errorf("failed to update sites: %v", err)
        helper.FailWithMessage("更新失败:"+err.Error(), c)
        return
    }

    c.JSON(http.StatusOK, gin.H{})
}

// CreateSite
// @Param data
// @Router /api/sites/create [post]
func CreateSite(c *gin.Context) {
    var req model.SiteCreateReq
    if err := c.ShouldBindJSON(&req); err != nil {
        logger.Errorf("CreateSite|ShouldBindJSON failed:%v", err)
        helper.FailWithMessage(fmt.Sprintf("ShouldBindJSON failed:%v", err), c)
        return
    }

    if err := model.CreateSite(req); err != nil {
        logger.Errorf("CreateSite model failed:%v", err)
        helper.FailWithMessage("创建失败:"+err.Error(), c)
        return
    }

    helper.OkWithMessage("创建成功", c)
    return
}

func SiteConfig(c *gin.Context) {
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

func UpdateSite(c *gin.Context) {

}
