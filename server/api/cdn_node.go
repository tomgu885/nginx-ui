package api

import (
    "github.com/gin-gonic/gin"
    "github.com/spf13/cast"
    "net/http"
    "nginx-ui/pkg/helper"
    "nginx-ui/server/model"
)

// GetCdnNodes
// Get /api/cdn_nodes
func GetCdnNodes(c *gin.Context) {
    data, err := model.GetCdnNodes(c)
    if err != nil {

        helper.FailWithMessage("获取数据错误"+err.Error(), c)
        return
    }

    c.JSON(http.StatusOK, data)
}

// GetCdnNode
// GET /api/cdn_nodes/1
func GetCdnNode(c *gin.Context) {
    id := cast.ToUint(c.Param("id"))
    node, err := model.GetCdnNodeById(id)
    if err != nil {
        helper.FailWithMessage("获取数据错误"+err.Error(), c)
        return
    }

    c.JSON(200, node)
}

// CreateCdnNode
// POST /api/cdn_nodes
func CreateCdnNode(c *gin.Context) {
    var req model.CdnNodeReq
    if err := c.ShouldBindJSON(&req); err != nil {
        helper.FailWithMessage("json绑定失败:"+err.Error(), c)
        return
    }
    err := model.CreateCdnNode(req)
    if err != nil {
        helper.FailWithMessage("创建失败:"+err.Error(), c)
        return
    }
    helper.OkWithMessage("创建成功", c)
    return
}

func UpdateCdnNode(c *gin.Context) {

}

// DeleteCdnNode
// DELETE /api/cdn_nodes/:id
func DeleteCdnNode(c *gin.Context) {
    id := cast.ToUint(c.Param("id"))
    _, err := model.GetCdnNodeById(id)
    if err != nil {
        helper.FailWithMessage("获取数据错误"+err.Error(), c)
        return
    }

    err = model.DeleteCdnNode(id)
    if err != nil {
        helper.FailWithMessage("删除错误"+err.Error(), c)
        return
    }

    helper.OkWithMessage("删除成功", c)
}
