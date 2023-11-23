package model

import "github.com/gin-gonic/gin"

type CdnNode struct {
    BaseModel
    Title  string `json:"title"`
    ApiUrl string `json:"api_url"`
    State  int8
}

func (CdnNode) TableName() string {
    return "cdn_nodes"
}

type CdnNodeReq struct {
    Title  string `json:"title"`
    ApiUrl string `json:"api_url"`
    State  int8   `json:"state"`
}

func GetCdnNodeById(id uint) (item CdnNode, err error) {
    err = db.Where("id", id).First(&item).Error
    return
}

func CreateCdnNode(req CdnNodeReq) (err error) {
    db.Create(&CdnNode{
        Title:  req.Title,
        ApiUrl: req.ApiUrl,
        State:  req.State,
    })
    return
}

func UpdateCdnNode(id uint, req CdnNodeReq) (err error) {
    err = db.Model(&CdnNode{}).Where("id", id).Updates(req).Error
    return
}

func DeleteCdnNode(id uint) (err error) {
    err = db.Delete(&CdnNode{}, id).Error
    return
}

func GetCdnNodes(c *gin.Context) (data DataList, err error) {
    var total int64
    q := db.Model(&CdnNode{})
    err = q.Count(&total).Error

    if err != nil {
        return
    }

    var list []CdnNode
    err = q.Find(&list).Error

    if err != nil {
        return
    }

    data = GetListWithPagination(&list, c, total)

    return
}
