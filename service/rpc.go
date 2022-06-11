package service

import (
	"SearchEngineV2/global"
	"SearchEngineV2/searcher"
)

type Hint struct {
	Container *searcher.Container
}

func NewHint() *Hint {
	return &Hint{
		Container: global.Container,
	}
}

// WordCut 分词
func (w *Hint) WordFindEnd(keyword string) []string {
	hints := w.Container.RpcEngine.FindEnd(keyword)
	return hints
}
