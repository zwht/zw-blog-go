package result

type Result struct {
	No  int         `json:"no"`
	Msg string      `json:"msg"`
	Obj interface{} `json:"obj,omitempty"`
}
