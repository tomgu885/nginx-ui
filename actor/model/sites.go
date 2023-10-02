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
    Result
    Data struct {
        List  []model.Site `json:"list"`
        Total uint         `json:"total"`
    } `json:"data"`
}

type UpdateReq struct {
    Force int `json:"force"` // 1 force, 2: not force

}
