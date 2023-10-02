package model

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

type LogCreateReq struct {
    NodeId string `json:"node_id"`
}
