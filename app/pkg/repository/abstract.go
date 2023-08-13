package repository

type PageModel struct {
	CurPage  uint        `json:"cur_page"`
	LastPage uint        `json:"last_page"`
	Total    uint64      `json:"total"`
	PageSize uint        `json:"page_size"`
	Data     interface{} `json:"data"`
}

type Abstract struct {
}
