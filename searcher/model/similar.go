package model

// SimilarhRequest 搜索请求
type SimilarhRequest struct {
	Query string `json:"query,omitempty" form:"database"` // 搜索关键词
}

// SimilarhResult 搜索响应
type SimilarhResult struct {
	Total     int           `json:"total"`               //总数
	Words     []string      `json:"words,omitempty"`     //搜索关键词
	Documents []ResponseDoc `json:"documents,omitempty"` //文档
}
