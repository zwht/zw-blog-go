package datamodels

type ResultPage struct {
	TotalCount int         `json:"totalCount"`
	PageSize   int         `json:"pageSize"`
	PageNum    int         `json:"pageNum"`
	PageData   interface{} `json:"pageData,omitempty"`
}
