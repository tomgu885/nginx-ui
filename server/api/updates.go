package api

import (
	"github.com/gin-gonic/gin"
	"nginx-ui/pkg/helper"
	"nginx-ui/server/model"
)

// cert & sites
func Updates(c *gin.Context) {
	// @todo sign, valid
	var req model.SiteListReq
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
