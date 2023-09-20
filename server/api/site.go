package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"nginx-ui/server/internal/helper"
	"nginx-ui/server/internal/logger"
	"nginx-ui/server/model"
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
// @Router /api/sites/create [post]
func CreateSite(c *gin.Context) {
	var req model.SiteCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.FailWithMessage(fmt.Sprintf("ShouldBindJSON failed:%v", err), c)
	}

}
