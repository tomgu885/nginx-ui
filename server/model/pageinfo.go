package model

type PageInfo struct {
	StartCreatedAt int64  `json:"start_created_at" form:"start_created_at"`
	EndCreatedAt   int64  `json:"end_created_at" form:"end_created_at"`
	Page           int    `json:"page" form:"page"`         // 页码
	PageSize       int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword        string `json:"keyword" form:"keyword"`   //关键字
	Ids            []uint `json:"ids" form:"ids"`
}

func (p PageInfo) GetPage() int {
	if p.Page <= 1 {
		return 1
	}

	return p.Page
}

func (p PageInfo) GetPageSize() int {
	if p.PageSize <= 0 {
		return 20
	}

	return p.PageSize
}

func (p PageInfo) GetOffsetLimit() (offset int, limit int) {
	limit = p.PageSize
	if limit <= 0 {
		limit = 20
	}

	page := p.Page

	if page <= 1 {
		page = 1
	}

	offset = limit * (page - 1)

	return
}
