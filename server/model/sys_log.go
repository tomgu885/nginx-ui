package model

import "github.com/gin-gonic/gin"

type SysLog struct {
    ID          int    `json:"id" gorm:"primaryKey"`
    NodeId      string `json:"node_id"`
    TraceId     string `json:"trace_id"`
    Restarted   int    `json:"restarted"`    // 是否重启 nginx 1:重启, 2:没有
    SiteUpdated uint   `json:"site_updated"` // 更新站点数量
    Content     string
    CreatedAt   int64 `json:"created_at" gorm:"autoCreateTime"`
}

func (SysLog) TableName() string {
    return "sys_log"
}

type SysLogListReq struct {
    PageInfo
    TraceId   string `json:"trace_id" form:"trace_id"`
    NodeId    string `json:"node_id" form:"node_id"`
    Restarted int8   `json:"restarted" form:"restarted"`
    Content   string `json:"content" form:"content"`
}

type LogCreateReq struct {
    NodeId      string `json:"node_id"`
    Content     string `json:"content"`
    Restarted   int    `json:"restarted"`
    SiteUpdated uint   `json:"site_updated"`
}

func CreateLog(req LogCreateReq, requestId string) (err error) {
    return db.Create(&SysLog{
        ID:          0,
        NodeId:      req.NodeId,
        TraceId:     requestId,
        Restarted:   req.Restarted,
        SiteUpdated: req.SiteUpdated,
        Content:     req.Content,
    }).Error
}

func GetSysLogs(c *gin.Context, req SysLogListReq) (data DataList, err error) {
    var total int64
    query := db.Model(&SysLog{})

    if len(req.Ids) > 0 {
        query.Where("id in ?", req.Ids)
    } else {
        if req.EndCreatedAt > 0 {
            query.Where("created_at <= ?", req.EndCreatedAt)
        }

        if req.StartCreatedAt > 0 {
            query.Where("created_at >= ?", req.StartCreatedAt)
        }

        if req.TraceId != "" {
            query.Where("trace_id", req.TraceId)
        }

        if req.NodeId != "" {
            query.Where("node_id", req.NodeId)
        }

        if req.Restarted != 0 {
            query.Where("restarted", req.Restarted)
        }

        if req.Content != "" {
            query.Where("content like ?", "%"+req.Content+"%")
        }
    }

    err = query.Count(&total).Error
    if err != nil {
        return
    }

    var list []SysLog
    offset, limit := req.GetOffsetLimit()
    err = query.Order("id DESC").Offset(offset).Limit(limit).Find(&list).Error
    if err != nil {
        return
    }
    data = GetListWithPagination(list, c, total)
    return
}
