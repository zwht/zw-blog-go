package datamodels

type ResultPage struct {
	TotalCount int         `json:"totalCount"`
	PageData   interface{} `json:"pageData,omitempty"`
}
