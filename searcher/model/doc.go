package model

type ResponseDoc struct {
	OriginalText string   `json:"originalText,omitempty"`
	Score        int      `json:"score,omitempty"` //得分
	Keys         []string `json:"keys,omitempty"`
	Url          string
	Loc          int
}
