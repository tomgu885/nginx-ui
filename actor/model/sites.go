package model

import "nginx-ui/server/model"

type Result struct {
    Code uint
    Data any
}

type PageResult struct {
    Result
    Data struct {
        List  any `json:"list"`
        Total uint
    } `json:"data"`
}

type SitePageResult struct {
    Code int8
    Data []model.Site `json:"data"`

    Pagination struct {
        Total       int64
        PerPage     int64 `json:"per_page"`
        CurrentPage int64 `json:"current_page"`
        TotalPage   int64 `json:"total_page"`
    }
}

type UpdateReq struct {
    Force   int `json:"force"`   // 1 force, 2: not force
    Restart int `json:"restart"` // 重启 1: restart , 2:reload

}
