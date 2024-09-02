package utils

type Response struct {
	Message    interface{} `json:"message"`
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
}
