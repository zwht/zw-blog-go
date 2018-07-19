package datamodels

type Result struct {
	ErrorCode int         `json:"errorCode"`
	Msg       string      `json:"msg"`
	Data      interface{} `json:"data,omitempty"`
}
