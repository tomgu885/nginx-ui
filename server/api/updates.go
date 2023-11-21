package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nginx-ui/pkg/helper"
	"nginx-ui/server/model"
)

func Report(c *gin.Context) {
	var req model.LogCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		helper.FailWithMessage("json错误:"+err.Error(), c)
		return
	}

	model.CreateLog(req, c.GetString("x-request-id"))
	helper.OkWithMessage("保存成功", c)
	return
}

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
	data, err := model.GetSites(c, req)
	//total := data.Pagination.Total
	if err != nil {
		helper.FailWithMessage("获取失败", c)
		return
	}

	c.JSON(http.StatusOK, data)
}
