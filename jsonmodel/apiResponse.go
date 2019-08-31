package jsonmodel

// APIResponse struct
type APIResponse struct {
	HasError bool        `json:"hasError"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}
