package model

type Response struct {
	Success bool
	Message string      `json:"Message"`
	Data    interface{} `json:"data,omitempty"`
}
