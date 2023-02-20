package api

type Format struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Errors  []any       `json:"errors"`
}
