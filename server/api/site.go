package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	list, total, err := model.GetSites(req)
	if err != nil {
		logger.Errorf("GetSites faieled: %v", err)
		helper.FailWithMessage("获取数据失败", c)
		return
	}

	helper.OkWithDetailed(helper.PageResult{
		Page:  req.Page,
		Total: total,
		List:  list,
	}, "获取成功", c)
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
