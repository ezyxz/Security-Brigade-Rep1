package model

type ResponseData struct {
	State   bool        `json:"state"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
