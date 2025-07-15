package response

type ListData[T any] struct {
	Total int64  `json:"total"`
	Rows  []T    `json:"rows"`
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
}
